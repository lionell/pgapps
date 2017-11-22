package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/lionell/pgapps/common"
	"github.com/lionell/pgapps/database"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "port to listen on")

func main() {
	flag.Parse()

	db := database.NewPostgres()
	err := db.Open()
	if err != nil {
		log.Fatalf("Can't open DB connection: %v", err)
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)
	router.Path("/query").Queries("q", "{q}").HandlerFunc(common.WrapDatabase(db, common.QueryHandler))

	log.Printf("Running server on port %s.", *port)
	http.ListenAndServe(":"+*port, router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./res/ajax.html")
}
