package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters

var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    fristThreeChar,
	"ftime": monthDayYear,
}

func fristThreeChar(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("2006-02-01")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.tpl"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}
	data := struct {
		Wisdom    []sage
		Transport []car
		Timenow   time.Time
	}{
		sages,
		cars,
		time.Now(),
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.tpl", data)
	if err != nil {
		log.Fatalln(err)
	}
}
