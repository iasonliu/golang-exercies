package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	read(conn)
	write(conn)
}

func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for i := 0; scanner.Scan(); i++ {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			fmt.Printf("Method --> %s\r\n", strings.Fields(ln)[0])
			fmt.Printf("URI --> %s\r\n", strings.Fields(ln)[1])
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
}

func write(conn net.Conn) {
	body := `<h1>HELLO HTTP GET</h1>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
