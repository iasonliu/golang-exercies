package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		// Receive experssion is just a value
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm Leaveing")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Expression to be sent can be any suiable value
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
