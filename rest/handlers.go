package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lionell/pgapps/database"
	"io/ioutil"
	"log"
	"net/http"
)

func SelectAllHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table, ok := vars["table"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'table' is not specified")
		return
	}
	query := fmt.Sprintf("select * from %s", table)
	log.Print(query)

	res, err := db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "query execution error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Fatalf("error while encoding json: %v", err)
	}
}

func SelectHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table, ok := vars["table"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'table' is not specified")
		return
	}
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'id' is not specified")
		return
	}
	query := fmt.Sprintf("select * from %s where id = %s", table, id)
	log.Printf(query)

	res, err := db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "query execution error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Fatalf("error encoding json: %v", err)
	}
}

func CreateHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table, ok := vars["table"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'table' is not specified")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v", err)
		return
	}
	if err = r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v", err)
		return
	}
	var recv database.Table
	if err := json.Unmarshal(body, &recv); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error decoding json: %v", err)
		return
	}

	var cols bytes.Buffer
	for i, c := range recv.Header {
		cols.WriteString(c)
		if i != len(recv.Header)-1 {
			cols.WriteString(", ")
		}
	}

	for _, r := range recv.Rows {
		var vals bytes.Buffer
		for i, v := range r {
			vals.WriteString(fmt.Sprintf("'%s'", v))
			if i != len(r)-1 {
				vals.WriteString(", ")
			}
		}
		query := fmt.Sprintf("insert into %s (%s) values (%s)", table, cols.String(), vals.String())
		log.Print(query)

		_, err := db.Exec(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error inserting row: %v", err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table, ok := vars["table"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'table' is not specified")
		return
	}
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'id' is not specified")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v", err)
		return
	}
	if err = r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v", err)
		return
	}
	var recv database.Table
	if err := json.Unmarshal(body, &recv); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error decoding json: %v", err)
		return
	}

	for _, r := range recv.Rows {
		var vals bytes.Buffer
		for i, v := range r {
			vals.WriteString(fmt.Sprintf("%s = '%s'", recv.Header[i], v))
			if i != len(r)-1 {
				vals.WriteString(", ")
			}
		}
		query := fmt.Sprintf("update %s set %s where id = %s", table, vals.String(), id)
		log.Print(query)

		_, err := db.Exec(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error updating row: %v", err)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func DeleteHandler(db database.Engine, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	table, ok := vars["table"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'table' is not specified")
		return
	}
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "parameter 'id' is not specified")
		return
	}
	query := fmt.Sprintf("delete from %s where id = %s", table, id)
	log.Print(query)

	_, err := db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "query execution error: %v", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
