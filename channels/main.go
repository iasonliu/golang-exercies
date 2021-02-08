package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://stackoverflow.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	if _, err := http.Get(link); err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
