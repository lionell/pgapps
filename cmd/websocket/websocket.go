package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/lionell/pgapps/websocket"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "port to listen on")

func main() {
	flag.Parse()

	h := websocket.NewHub()
	go h.Run()

	s := websocket.NewSolver(h)
	go s.Run()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWebSocket(h, w, r)
	})

	log.Printf("Running server on port %s.", *port)
	http.ListenAndServe(":"+*port, router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./res/websocket.html")
}
