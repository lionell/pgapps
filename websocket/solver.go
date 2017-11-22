package websocket

import (
	"github.com/lionell/pgapps/database"
	"github.com/lionell/pgapps/message"
	"log"
)

type Solver struct {
	hub *Hub
	db  database.Engine
}

func NewSolver(h *Hub) *Solver {
	db := database.NewPostgres()
	db.Open()
	return &Solver{
		db:  db,
		hub: h,
	}
}

func (s Solver) Run() {
	for {
		select {
		case query := <-s.hub.Queries:
			log.Printf("Executing query '%s'.", query)
			res, err := s.db.Exec(query)
			if err != nil {
				s.hub.Broadcast <- &message.Message{Query: query, Error: err.Error()}
				break
			}
			s.hub.Broadcast <- &message.Message{Query: query, Result: *res}
		}
	}
}
