package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password []byte
}

type Users []User

var UserList Users

const mySigningKey = "AllYOURBASEKey"

const MyHMACkey = "my secret key 001 james bond rule the world"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func getUser(email string) (User, error) {
	for _, u := range UserList {
		if u.Email == email {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User not Find")
}

func index(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("errormsg")

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Web Example</title>
	</head>
	<body>
	<h1>IF THERE IS ANY ERROR: ` + errMsg + `</h1>
	<div>
	<h1> REGISTER </h1>
	<form action="/register" method="post">
		<input type="email" name="emailForm"/>
		<input type="password" name="passwordForm"/>
		<input type="submit" name="register"/>
	</form>
	</div>
	<div>
	<h1> LOGIN IN <h1>
	<form action="/login" method="post">
		<input type="email" name="emailForm"/>
		<input type="password" name="passwordForm"/>
		<input type="submit" name="Login"/>
	</form>
	</div>
	</body>
	</html>`
	io.WriteString(w, html)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userEmail := r.FormValue("emailForm")
		userPassword := r.FormValue("passwordForm")
		if userEmail == "" || userPassword == "" {
			errorMsg := url.QueryEscape("Your Email or Password empty")
			http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
		}
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
		if err != nil {
			errorMsg := url.QueryEscape("there was Internal server Error!!")
			http.Error(w, errorMsg, http.StatusInternalServerError)
		}
		UserList = append(UserList, User{userEmail, hashPassword})
		fmt.Fprintf(w, "%#v", UserList)
	} else {
		errorMsg := url.QueryEscape("Your method was not Post")
		http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userEmail := r.FormValue("emailForm")
		userPassword := r.FormValue("passwordForm")
		if userEmail == "" || userPassword == "" {
			errorMsg := url.QueryEscape("Your Email or Password empty")
			http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
			return
		}
		u, err := getUser(userEmail)
		if err != nil {
			errorMsg := url.QueryEscape("User Not find")
			http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
			return
		}
		err = bcrypt.CompareHashAndPassword(u.Password, []byte(userPassword))
		if err != nil {
			errorMsg := url.QueryEscape("Username or Password error")
			http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
			return
		}
		fmt.Fprintf(w, "Login %s", u.Email)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createToken(sessionId string) string {
	mac := hmac.New(sha256.New, []byte(MyHMACkey))
	mac.Write([]byte(sessionId))
	// to hex
	// signedMac := fmt.Sprintf("%x", mac.Sum(nil))
	// to base64
	signedMacBase64 := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signedMacBase64 + "|" + sessionId
}

func paresToken(ss string) (string, error) {
	xs := strings.SplitN(ss, "|", 2)
	if len(xs) != 2 {
		return "", fmt.Errorf("Session id not Signed")
	}
	signedMac, err := base64.StdEncoding.DecodeString(xs[0])
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, []byte(MyHMACkey))
	mac.Write([]byte(xs[1]))
	if hmac.Equal([]byte(signedMac), mac.Sum(nil)) {
		return xs[1], nil
	}
	return "", fmt.Errorf("Session id not Signed")
}
