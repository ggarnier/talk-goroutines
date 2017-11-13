package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Results struct {
	okCount  int
	errCount int
	lock     *sync.Mutex
}

func runChecks(urls []string) Results {
	r := Results{lock: &sync.Mutex{}}
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Head(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				r.lock.Lock()
				r.errCount++
				r.lock.Unlock()
			} else {
				r.lock.Lock()
				r.okCount++
				r.lock.Unlock()
			}
		}(url)
	}

	// Wrong way to wait for goroutines to finish
	time.Sleep(2 * time.Second)

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
