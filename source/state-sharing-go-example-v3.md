### Version 3: using channels

```go
	func runChecks(urls []string) Results {
		r := Results{}
		responses := make(chan bool, len(urls))
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
```

<span class="fragment current-only" data-code-focus="3,8"></span>
<span class="fragment current-only" data-code-focus="13"></span>

<!--
- no need to use locks
- less code
- better performance
-->
