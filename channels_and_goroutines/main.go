package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://en.wikipedia.org/wiki/One_Piece",
		"https://www.amazon.com/",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for receivedresponse := range channel {

		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, channel)
		}(receivedresponse)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Might be down")
		fmt.Println(link, "is down")
		channel <- link
	} else {
		fmt.Println(link, "is up")
		fmt.Println("Yep it's up")
		channel <- link
	}
}
