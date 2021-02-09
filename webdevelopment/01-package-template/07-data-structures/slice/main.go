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
	person := []string{"tom", "tim", "jason"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.tpl", person)
	if err != nil {
		log.Fatalln(err)
	}
}
