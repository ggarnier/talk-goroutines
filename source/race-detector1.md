## Race Detector tool

```go
	func runChecks(urls []string) Results {
		r := Results{}
		lock := sync.Mutex{}
		for _, url := range urls {
			go func(url string) {
				resp, err := http.Head(url)
				if err != nil || resp.StatusCode != http.StatusOK {
					lock.Lock()
					r.errCount++
					lock.Unlock()
				} else {
					lock.Lock()
					r.okCount++
					lock.Unlock()
				}
			}(url)
		}

		// Wrong way to wait for goroutines to finish
		time.Sleep(2 * time.Second)

		return r
	}
```

<span class="fragment current-only" data-code-focus="3,8,10,12,14"></span>
<span class="fragment current-only" data-code-focus="19-20"></span>
