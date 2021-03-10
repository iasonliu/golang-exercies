package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func greet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		id := uuid.Must(uuid.NewRandom())
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		})
	}
	fmt.Println(cookie.Value)

}

func main() {
	http.HandleFunc("/", greet)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
