package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		// open
		file, fileHeader, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// for your infomation
		fmt.Println("\nfile:", file, "\nfile header:", fileHeader, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		s = string(bs)

		// store file into path
		err = ioutil.WriteFile(filepath.Join("./file", fileHeader.Filename), bs, 0666)
		// dst, err := os.Create(filepath.Join("./file", fileHeader.Filename))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	return
		// }
		// defer dst.Close()

		// _, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
