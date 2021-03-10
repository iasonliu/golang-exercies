package main

import "net/http"

func getUser(r *http.Request) user {
	var u user
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}
	// if the user exists already, get user
	if username, ok := dbSessions[c.Value]; ok {
		u = dbUsers[username]
	}
	return u
}

func isLogged(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	username := dbSessions[c.Value]
	_, ok := dbUsers[username]
	return ok
}
