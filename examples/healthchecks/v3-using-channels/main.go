package main

import (
	"fmt"
	"net/http"
)

type Results struct {
	okCount  int
	errCount int
}

func runChecks(urls []string) Results {
	r := Results{}
	responses := make(chan bool)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Head(url)
			success := err == nil && resp.StatusCode == http.StatusOK
			responses <- success
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		success := <-responses
		if success {
			r.okCount++
		} else {
			r.errCount++
		}
	}

	return r
}

func main() {
	urls := []string{
		"http://yahoo.com",
		"https://www.yahoo.com",
		"http://wrong-url.com",
	}

	r := runChecks(urls)
	fmt.Printf("%d ok, %d errors\n", r.okCount, r.errCount)
}
