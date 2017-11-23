package main

import (
	"github.com/gorilla/mux"
	"github.com/lionell/pgapps/database"
	"github.com/lionell/pgapps/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	db := database.NewPostgres()
	err := db.OpenRemote(os.Getenv("MY_POSTGRES_SERVICE_HOST"), os.Getenv("MY_POSTGRES_SERVICE_PORT"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	h := websocket.NewHub()
	go h.Run()

	s := websocket.Solver{
		Db:  db,
		Hub: h,
	}
	go s.Run()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWebSocket(h, w, r)
	})

	log.Print("Running server on port 8080.")
	http.ListenAndServe(":8080", router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./websocket.html")
}
