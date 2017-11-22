package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/lionell/pgapps/message"
	"log"
	"net/http"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	Send chan *message.Message
}

func (c *Client) runReadListener() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Fatal(err)
			}
			break
		}
		c.hub.Queries <- string(msg)
	}
}

func (c *Client) runWriteListener() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteJSON(msg)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWebSocket(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := &Client{hub: h, conn: conn, Send: make(chan *message.Message)}
	h.Register <- client

	go client.runWriteListener()
	go client.runReadListener()
}
