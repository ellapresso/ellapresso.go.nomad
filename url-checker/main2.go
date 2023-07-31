package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main2() {
	// // example 01
	// var results map[string]string
	// results["hello"] = "Hello"
	// // 위 경우 panic 발생한다 - 초기화 되지 않은 map을 사용하려고 했기 때문.

	// // example 02
	// var results = map[string]string{}
	// results["hello"] = "Hello"
	// // 정상 동작함

	// // example 03
	// var results = make(map[string]string)
	// results["hello"] = "Hello"
	// fmt.Println(results)
	// // 정상 동작함

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

	// for _, url := range urls {
	// 	fmt.Println(url)
	// }

	var results = make(map[string]string)
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
