package main

import (
	"flag"
	"github.com/lionell/pgapps/database"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	eng database.Engine
}

func (s *Server) Exec(query string, result *database.Table) error {
	log.Printf("Executing query '%s'", query)
	t, err := s.eng.Exec(query)
	if err != nil {
		return err
	}

	result.Header = t.Header
	result.Rows = t.Rows
	return nil
}

var port = flag.String("port", "1234", "port to listen on")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Can't listen on port 1234")
		return
	}
	log.Printf("Listening on port %s", *port)

	p := database.NewPostgres()
	err = p.Open()
	if err != nil {
		log.Fatalf("Can't open DB connection: %v", err)
	}
	defer p.Close()

	s := &Server{p}
	rpc.Register(s)
	rpc.Accept(l)
}
