package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	msg := "Hello World!!!"
	h := sha512.New()
	h.Write([]byte(msg))
	fmt.Printf("msg: %x\n", h.Sum(nil))

	// sha with file
	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// defer f.Close()
	// if _, err := io.Copy(h, f); err != nil {
	// 	log.Fatalln(err)
	// }

	bs, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	h = sha256.New()
	h.Write(bs)

	fmt.Printf("here's the type BEFORE Sum: %T\n", h)
	fmt.Printf("here's BEFORE Sum: %#v\n", h)
	fmt.Printf("file: %x\n", h.Sum(nil))
}
