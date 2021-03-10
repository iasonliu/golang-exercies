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

func init() {
	db, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/testdb?charset=utf8")
	check(err)
	err = db.Ping()
	check(err)
}
func main() {
	defer db.Close()
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, _ *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	rs, err := stmt.Exec()
	check(err)

	n, err := rs.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, _ *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	rs, err := stmt.Exec()
	check(err)

	n, err := rs.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD INTO customer", n)
}

func read(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)

	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, _ *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	rs, err := stmt.Exec()
	check(err)

	n, err := rs.RowsAffected()
	fmt.Fprintln(w, "UPDATE RECORD", n)
}

func del(w http.ResponseWriter, _ *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
