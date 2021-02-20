package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s; and the request method is %s on %s\n", time.Now(), r.Method, r.URL)
	fmt.Printf("Hello World! %s; and the request method is %s on %s\n", time.Now(), r.Method, r.URL)
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("the request method is %s on %s\n", r.Method, r.URL)
	// process form data
	w.Header().Set("Location", "/")
	// 301 -- StatusMovedPermanently
	// w.WriteHeader(http.StatusMovedPermanently)
	// 303 --StatusSeeOther
	// w.WriteHeader(http.StatusSeeOther)
	// 307 -- StatusTemporaryRedirect
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("the request method is %s on %s\n", r.Method, r.URL)
	tpl.ExecuteTemplate(w, "index.tpl", nil)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/post", post)
	http.HandleFunc("/index", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
