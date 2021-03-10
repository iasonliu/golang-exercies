package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var (
		even  int
		odd   int
		total int
	)

	numbers := []int{1, 2, 3, 4, 0, 7, 8, 9, 10, -1}
	var err error
	for _, n := range numbers {
		total += 1
		switch {
		case n == 0:
		case n%2 == 0:
			even += 1
		case n%2 == 1:
			odd += 1
		default:
			err = errors.New("Found negative number")
			break
		}
	}
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Even %d, odd %d, total %d\n", even, odd, total)
}
