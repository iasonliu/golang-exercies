package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("pls add args as filename")
	}
	filename := os.Args[1]
	readfile(filename)
}

func readfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
}
