package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// add route to serve pictures
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("filename")
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
		defer file.Close()
		// create sha for file name
		h := sha1.New()
		if _, err := io.Copy(h, file); err != nil {
			log.Fatal(err)
		}
		uFilename := fmt.Sprintf("%x.%s", h.Sum(nil), strings.Split(fileHeader.Filename, ".")[1])
		// upload file, seek file
		file.Seek(0, 0)
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
		err = ioutil.WriteFile(filepath.Join("./public/pics/", uFilename), bs, 0755)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
		// add filename to this user's cookie
		cookie = appendValue(w, cookie, uFilename)
	}
	data := strings.Split(cookie.Value, "|")
	// sliced cookie values to only send over images
	tpl.ExecuteTemplate(w, "index.tpl", data[1:])
}

// add func to get cookie
func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		sID := uuid.Must(uuid.NewRandom())
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendValue(w http.ResponseWriter, cookie *http.Cookie, filename string) *http.Cookie {
	//append
	if !strings.Contains(cookie.Value, filename) {
		cookie.Value += fmt.Sprintf("|%s", filename)
	}
	http.SetCookie(w, cookie)
	return cookie
}
