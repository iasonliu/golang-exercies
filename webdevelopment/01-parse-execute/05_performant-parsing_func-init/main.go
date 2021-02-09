package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "three.tpl", nil); err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.tpl", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.tpl", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
