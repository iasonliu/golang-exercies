package main

import "fmt"

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range n {
		if v%2 == 0 {
			fmt.Printf("%v is even \n", v)
		} else {
			fmt.Printf("%v is odd \n", v)
		}
	}
}
