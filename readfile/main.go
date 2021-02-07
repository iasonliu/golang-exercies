package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("pls add args as filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	readfile(filename)
}

func readfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
