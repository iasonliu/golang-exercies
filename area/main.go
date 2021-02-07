package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func main() {
	sq := square{sideLength: 10.0}
	printArea(sq)
	tr := triangle{height: 10, base: 21.0}
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
