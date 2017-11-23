package websocket

import (
	"github.com/lionell/pgapps/database"
	"github.com/lionell/pgapps/message"
	"log"
)

type Solver struct {
	Hub *Hub
	Db  database.Engine
}

func NewSolver(h *Hub) *Solver {
	db := database.NewPostgres()
	db.Open()
	return &Solver{
		Db:  db,
		Hub: h,
	}
}

func (s Solver) Run() {
	for {
		select {
		case query := <-s.Hub.Queries:
			log.Printf("Executing query '%s'.", query)
			res, err := s.Db.Exec(query)
			if err != nil {
				s.Hub.Broadcast <- &message.Message{Query: query, Error: err.Error()}
				break
			}
			s.Hub.Broadcast <- &message.Message{Query: query, Result: *res}
		}
	}
}
