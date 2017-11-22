package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lionell/pgapps/database"
	"net/rpc"
	"os"
)

var (
	host = flag.String("host", "localhost", "address of the server")
	port = flag.String("port", "1234", "port to connect to")
)

func main() {
	flag.Parse()

	client, err := rpc.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Printf("Can't open connection: %v", err)
		return
	}
	fmt.Printf("Connection to %s:%s established.\n", *host, *port)

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()
		if s.Text() == "exit" {
			fmt.Println("Bye!")
			break
		}
		var result database.Table
		err = client.Call("Server.Exec", s.Text(), &result)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%v\n", result)
	}
}
