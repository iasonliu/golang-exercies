package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", post)
	http.HandleFunc("/url", url)
	http.HandleFunc("/get", get)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func post(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	<input type="text" name="q">
	<input type="submit"></form><br>
	`+v)
}

// http://localhost:8080/?getfromurl=dog
func url(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("getfromurl")
	io.WriteString(w, "vaule get from url is :"+data)
}

func get(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("getvaule")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	<input type="text" name="getvaule">
	<input type="submit"></form><br>
	`+v)
}
