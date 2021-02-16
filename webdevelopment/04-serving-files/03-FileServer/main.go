package main

import (
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)

}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
