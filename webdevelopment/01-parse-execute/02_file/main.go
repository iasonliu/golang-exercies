package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	defer newFile.Close()

	err = tpl.Execute(newFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
