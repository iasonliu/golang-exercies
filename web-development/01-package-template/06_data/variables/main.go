package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.tpl"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.tpl", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
