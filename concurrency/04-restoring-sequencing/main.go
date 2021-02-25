package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	text string
	wait chan bool
}

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		msg1 := <-ch
		fmt.Println(msg1.text)
		msg2 := <-ch
		fmt.Println(msg2.text)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring; I'm leaving.")
}
func fanIn(ch1, ch2 <-chan Message) <-chan Message {
	new_ch := make(chan Message)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
}

func generator(msg string) chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	// launch the goroutine form inside the func
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}
