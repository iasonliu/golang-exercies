package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I felt so good like anything was possible\n Ihit \n last line \n"
	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
