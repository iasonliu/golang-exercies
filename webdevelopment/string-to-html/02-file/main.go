package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println("os args[0] is ", os.Args[0])
	fmt.Println("os args[1] is ", os.Args[1])
	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>
`
	// err := ioutil.WriteFile("index.html", []byte(tpl), 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error createing file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
