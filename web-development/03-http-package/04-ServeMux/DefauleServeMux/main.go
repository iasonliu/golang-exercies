package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog dog!!!")

}

type hotcat int

func (hc hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat cat !!!")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)
	http.ListenAndServe(":8080", nil)
}
