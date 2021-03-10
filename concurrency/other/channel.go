package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := greet("Joe")
	for i := 0; i < 15; i++ {
		fmt.Printf("%q\n", <-ch)
	}
	for i := 0; i < 20; i++ {
		fmt.Println(rand.Intn(10))
	}
}

func greet(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s---> %d", msg, i)
		}
	}()
	return ch
}
