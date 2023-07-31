package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type requesResult struct {
	url    string
	status string
}

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	results := make(map[string]string)
	c := make(chan requesResult)
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requesResult) { // 이 채널은 보내기만 하겠다! send Only
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} else {
		c <- requesResult{url: url, status: status}
	}

}
