package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("test-key", "this is from test")
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "AtEnd3")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
