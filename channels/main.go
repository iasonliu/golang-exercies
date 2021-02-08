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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	if _, err := http.Get(link); err != nil {
		fmt.Println(link, "might be down!")
		c <- "Migth be down I think" + link
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep it's up! " + link
}
