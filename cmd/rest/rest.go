package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/lionell/pgapps/common"
	"github.com/lionell/pgapps/database"
	"github.com/lionell/pgapps/rest"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "port to listen on")

func main() {
	flag.Parse()

	db := database.NewPostgres()
	err := db.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/{table}").HandlerFunc(common.WrapDatabase(db, rest.SelectAllHandler))
	router.Methods("GET").Path("/{table}/{id}").HandlerFunc(common.WrapDatabase(db, rest.SelectHandler))
	router.Methods("POST").Path("/{table}").HandlerFunc(common.WrapDatabase(db, rest.CreateHandler))
	router.Methods("PUT").Path("/{table}/{id}").HandlerFunc(common.WrapDatabase(db, rest.UpdateHandler))
	router.Methods("DELETE").Path("/{table}/{id}").HandlerFunc(common.WrapDatabase(db, rest.DeleteHandler))

	log.Printf("Running server on port %s.", *port)
	http.ListenAndServe(":"+*port, router)
}
