package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)

	http.ListenAndServe(":8080", router)
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s\n", ps.ByName("article"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.tmpl", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.tmpl", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.tmpl", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.tmpl", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.tmpl", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
