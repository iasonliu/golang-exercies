package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func greet(w http.ResponseWriter, _ *http.Request) {
	_, err = io.WriteString(w, "Successfully complated.")
	check(err)
}

func main() {
	// user:password@tcp(localhost:555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/testdb?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
