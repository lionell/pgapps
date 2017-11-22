package websocket

import (
	"github.com/lionell/pgapps/message"
	"log"
)

type Hub struct {
	clients    map[*Client]bool
	Broadcast  chan *message.Message
	Queries    chan string
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Broadcast:  make(chan *message.Message),
		Queries:    make(chan string),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			log.Print("Client connected.")
			h.clients[c] = true
		case c := <-h.Unregister:
			log.Print("Client disconnected.")
			delete(h.clients, c)
			close(c.Send)
		case msg := <-h.Broadcast:
			for c := range h.clients {
				c.Send <- msg
			}
		}
	}
}
