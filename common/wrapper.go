package common

import (
	"github.com/lionell/pgapps/database"
	"net/http"
)

func WrapDatabase(db database.Engine, fn func(database.Engine, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(db, w, r)
	}
}
