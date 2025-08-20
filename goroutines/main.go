package main

import (
	"fmt"
	"net/http"
)

func checkUrl(url string, ch chan<- string) {
	_, err := http.Get(url)
	if err != nil {
		ch <- url + "is down"
		return
	}
	ch <- url + "is up !"
}

func main() {

	urls := []string{
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
	}

	ch := make(chan string, len(urls))
	for _, url := range urls {
		go checkUrl(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
