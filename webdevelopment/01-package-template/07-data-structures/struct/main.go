package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("index.tpl"))
}

func main() {
	jason := person{Name: "jason", Age: 12}
	err := tpl.ExecuteTemplate(os.Stdout, "index.tpl", jason)
	if err != nil {
		log.Fatalln(err)
	}
}
