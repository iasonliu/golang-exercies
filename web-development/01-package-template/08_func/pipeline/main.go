package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.tpl"))
}

func main() {
	data := 3
	err := tpl.ExecuteTemplate(os.Stdout, "index.tpl", data)
	if err != nil {
		log.Fatalln(err)
	}
}
