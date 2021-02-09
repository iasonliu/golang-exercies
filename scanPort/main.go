package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", port)
		}(i)
	}
	wg.Wait()
}
