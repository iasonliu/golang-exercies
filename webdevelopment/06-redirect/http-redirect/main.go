package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URI %s --> %s\n", r.URL, r.Method)
	http.Redirect(w, r, "/hello", http.StatusTemporaryRedirect)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URI %s --> %s\n", r.URL, r.Method)
	fmt.Fprintf(w, "Hello!!!!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
