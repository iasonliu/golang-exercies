package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first")
	lastName := r.FormValue("last")
	subscribe := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.tpl", person{firstName, lastName, subscribe})
	if err != nil {
		http.Error(w, "template error!!!", 500)
		log.Fatalln(err)
	}
}
