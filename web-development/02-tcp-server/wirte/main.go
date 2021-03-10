package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\n Hello from TCP server\n")
		fmt.Fprintf(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "How is your day?")
		conn.Close()
	}
}
