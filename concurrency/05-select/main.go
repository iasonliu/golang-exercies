package main

import "fmt"

type Message struct {
	text string
	wait chan bool
}

func main() {
	ch := fanIn(greet("Joe"), greet("Ann"))
	for i := 0; i < 15; i++ {
		msg1 := <-ch
		fmt.Println(msg1.text)
		msg2 := <-ch
		fmt.Println(msg2.text)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("ALL DONE!!!!!!!")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	ch := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				ch <- s
			case s := <-input2:
				ch <- s
			}
		}
	}()
	return ch
}

func greet(msg string) <-chan Message {
	ch := make(chan Message)
	waitForit := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s --> %d", msg, i), waitForit}
			<-waitForit
		}
	}()
	return ch
}
