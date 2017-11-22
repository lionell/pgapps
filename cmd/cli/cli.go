package main

import (
	"bufio"
	"fmt"
	"github.com/lionell/pgapps/database"
	"log"
	"os"
)

func main() {
	db := database.NewPostgres()
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()
		if s.Text() == "exit" {
			fmt.Println("Bye!")
			break
		}
		query := s.Text()
		table, err := db.Exec(query)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%v\n", table)
	}
}
