package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := fanIn(greet("Ann"), greet("Joe"))
	for i := 0; i < 15; i++ {
		fmt.Printf("%q\n", <-ch)
	}
	fmt.Printf("DONE!!!")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func greet(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s-%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
