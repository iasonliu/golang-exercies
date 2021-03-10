package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getCode(msg string) string {
	h := hmac.New(sha256.New, []byte("secret-key"))
	_, err := h.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	code := getCode(email)
	cookie := http.Cookie{
		Name:  "session",
		Value: code + "|" + email,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}
	isEqual := false
	xs := strings.SplitN(c.Value, "|", 2)
	if len(xs) == 2 {
		cCode := xs[0]
		cEmail := xs[1]
		code := getCode(cEmail)
		isEqual = hmac.Equal([]byte(cCode), []byte(code))
	}
	message := "Not logged in"
	if isEqual {
		message = "Logged in"
	}
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>HMAC Example</title>
	</head>
	<body>
		<p> Cookie value: ` + c.Value + `</p>
		<p>` + message + `</p>
	<form action="/submit" method="post">
		<input type="email" name="email"/>
		<input type="submit" />
	</form>
	</body>
	</html>`
	io.WriteString(w, html)
}
