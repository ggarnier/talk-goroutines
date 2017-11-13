## Race Detector tool

```go
	func runChecks(urls []string) Results {
		r := Results{lock: &sync.Mutex{}}
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
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

		wg.Wait() // Blocks until all `wg.Done()` calls
		return r
	}
```

<span class="fragment current-only" data-code-focus="3,5,7,21"></span>
