package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		// we don't checkLink to wait 5 seconds on the first call
		// so we made a function literal (lambda function) that
		// adds this delay.
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(<-c)
		// we pass <-c as a value so that the new goroutine
		// doesn't reference something that can change.
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		c <- link
		return
	}
	fmt.Println(link, "Is up!")
	c <- link
}
