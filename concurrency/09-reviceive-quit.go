package main

//  Deterministically quit goroutine with quit channel option in select

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	ch := greet("Hi!", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye!"
	fmt.Printf("Greet says %s", <-quit)
}

func greet(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				quit <- "See You!!"
				return
			}
		}
	}()
	return ch
}
