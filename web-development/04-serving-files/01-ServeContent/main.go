package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)

}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fileInfo.ModTime(), f)

}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	// Create sample handler to returns 404
	http.Handle("/resources", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
