package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.tmpl", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.tmpl"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
