package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func greet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("mycookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "mycookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

func main() {
	http.HandleFunc("/", greet)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
