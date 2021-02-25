package main

import (
	"fmt"
	"time"
)

func main() {
	ch := greet("Joe")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			// time.After returns a channel that waits N time to send a message
			fmt.Println("waited too long!")
			return
		}
	}
}

// return receive-only channel
func greet(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s --> %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
