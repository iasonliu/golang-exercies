package main

import "fmt"

type englishBot struct {
	name string
}

type spanishBot struct {
	name string
}

func (eb *englishBot) getGreeting() string {
	return "Hello " + eb.name + "!"
}

func (sb *spanishBot) getGreeting() string {
	return "Hello " + sb.name + "!"
}

func ebPrintGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func sbPrintGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{name: "eb"}
	ebPrintGreeting(eb)
	sb := spanishBot{name: "sb"}
	sbPrintGreeting(sb)
}
