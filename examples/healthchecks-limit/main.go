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
}

func runChecks(urls <-chan string) Results {
	max := 3 // Max parallel goroutines
	running := make(chan bool, max)
	r := Results{}
	responses := make(chan bool)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for url := range urls {
			wg.Add(1)
			running <- true
			go func(url string) {
				defer func() {
					<-running
					wg.Done()
				}()
				resp, err := http.Head(url)
				success := err == nil && resp.StatusCode == http.StatusOK
				responses <- success
			}(url)
		}
	}()

	go func() {
		for range time.Tick(100 * time.Millisecond) {
			fmt.Printf("%d goroutines running\n", len(running))
		}
	}()

	go func() {
		defer close(responses)
		wg.Wait()
	}()

	for success := range responses {
		if success {
			r.okCount++
		} else {
			r.errCount++
		}
	}

	return r
}

func main() {
	inputStream := func() <-chan string {
		urls := []string{
			"http://yahoo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://globo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://yahoo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://globo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://yahoo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://wrong-url.com/healthcheck",
			"http://globo.com",
			"http://xavier.globoplay.globoi.com/healthcheck",
			"http://xavier.globoplay.globoi.com/healthcheck",
		}
		stream := make(chan string)
		go func() {
			defer close(stream)
			for _, url := range urls {
				stream <- url
			}
		}()
		return stream
	}

	r := runChecks(inputStream())
	fmt.Printf("%d ok, %d errors\n", r.okCount, r.errCount)
}
