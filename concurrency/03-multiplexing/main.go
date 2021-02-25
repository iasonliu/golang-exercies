package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 15; i++ {
		fmt.Printf("%q\n", <-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	// launch the goroutine form inside the func
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s-->%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
