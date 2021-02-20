package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.tpl", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !isLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.tpl", u)
}
func list(w http.ResponseWriter, req *http.Request) {
	if isLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "listuser.tpl", dbUsers)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if isLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		// username taken?
		if _, ok := dbUsers[username]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.Must(uuid.NewRandom())
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		})
		dbSessions[sID.String()] = username

		// store user in dbUsers
		u := user{username, password, firstname, lastname}
		dbUsers[username] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.tpl", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if isLogged(req) {
		c, _ := req.Cookie("session")
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
