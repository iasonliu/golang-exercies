package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

const mySigningKey = "ALLYOURKEYSforsecrests"

func getJWT(msg string) (string, error) {

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Microsecond).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailForm")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tokenString, err := getJWT(email)
	if err != nil {
		http.Error(w, "coludn't getJWT", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:  "session",
		Value: tokenString,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}
	tokenString := c.Value
	afterVerificationToken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(beforeVeritificationToken *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	// StandardClaims has the Valid() error method
	// which mean it implements the Claims interface..
	// type Claims interface {
	// 	Valid() error
	// }
	// When you ParseWithClaims the Valid() method gets run
	// and if all is well, then return no "error" and
	// Type Token which has a field VALID will be true

	message := "Not logged in"
	if err == nil && afterVerificationToken.Valid {
		claims := afterVerificationToken.Claims.(*MyClaims)
		log.Println(claims.Email, claims.ExpiresAt)
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
		<input type="email" name="emailForm"/>
		<input type="submit" />
	</form>
	</body>
	</html>`
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}
