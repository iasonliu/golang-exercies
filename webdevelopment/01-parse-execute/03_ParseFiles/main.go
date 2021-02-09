package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.tpl")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.tpl", "three.tpl")
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
