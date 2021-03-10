package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "index")
}

func me(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "me!!!!")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
