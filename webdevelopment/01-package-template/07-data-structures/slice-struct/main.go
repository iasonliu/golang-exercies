package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("index-var.tpl"))
}

func main() {
	jason := person{Name: "jason", Age: 12}
	tim := person{Name: "tim", Age: 10}
	tom := person{Name: "tom", Age: 18}
	lucy := person{Name: "lucy", Age: 19}

	myfriends := []person{jason, tim, tom, lucy}

	err := tpl.ExecuteTemplate(os.Stdout, "index-var.tpl", myfriends)
	if err != nil {
		log.Fatalln(err)
	}
}
