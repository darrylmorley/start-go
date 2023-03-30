package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	l := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, u := range l {
		go checkLink(u, c)
	}

	for m := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(m)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l + " might be down!\n")
		c <- l
		return
	}

	fmt.Println(l + " is up!\n")
	c <- l
}
