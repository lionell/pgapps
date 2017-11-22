package common

import (
	"encoding/json"
	"github.com/lionell/pgapps/database"
	"github.com/lionell/pgapps/message"
	"log"
	"net/http"
)

func QueryHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	res, err := db.Exec(query)
	var msg *message.Message
	if err != nil {
		msg = &message.Message{Query: query, Error: err.Error()}
	} else {
		msg = &message.Message{Query: query, Result: *res}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		log.Fatalf("error encoding json: %v", err)
	}
}
