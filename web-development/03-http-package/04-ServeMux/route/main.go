package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "<h1>Dog Dog Dog !!!</h1>")
	case "/cat":
		io.WriteString(w, "<h1>Cat Cat Cat !!!</h1>")
	default:
		io.WriteString(w, "<h1>index !!!</h1>")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
