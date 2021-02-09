package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.tmpl", nil)
	if err != nil {
		log.Fatalln("template didn't execute:", err)
	}
}
