package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, _ *http.Request) {
	data := "This is from dog"
	tpl.ExecuteTemplate(w, "dog.gohtml", data)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpeg", dogPic)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
