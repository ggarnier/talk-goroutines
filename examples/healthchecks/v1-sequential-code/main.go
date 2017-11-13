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
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			r.errCount++
		} else {
			r.okCount++
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
