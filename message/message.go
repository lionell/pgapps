package message

import "github.com/lionell/pgapps/database"

type Message struct {
	Query  string         `json:"query"`
	Result database.Table `json:"result"`
	Error  string         `json:"error"`
}
