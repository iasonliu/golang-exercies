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

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "templates files error", http.StatusServiceUnavailable)
	}
}
