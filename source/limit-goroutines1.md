### Limiting the number of running goroutines

<!-- examples/healthchecks-limit/main.go -->

```go
	max := 3 // Max simultaneous goroutines
	running := make(chan struct{}, max)

	for url := range urls {
		running <- struct{}{} // waits for a free slot
		go func(url string) {
			defer func() {
				<-running // releases slot
			}()
			// do work
		}(url)
	}
```

<span class="fragment current-only" data-code-focus="1"></span>
<span class="fragment current-only" data-code-focus="2"></span>
<span class="fragment current-only" data-code-focus="5"></span>
<span class="fragment current-only" data-code-focus="8"></span>
