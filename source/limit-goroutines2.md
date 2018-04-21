### Limiting the number of running goroutines

<!-- examples/healthchecks-limit/main.go -->

```go
	max := 3 // Max simultaneous goroutines
	running := make(chan bool, max)

	go func() {
		for range time.Tick(100 * time.Millisecond) {
			fmt.Printf("%d goroutines running\n", len(running))
		}
	}()

	for url := range urls {
		running <- true // waits for a free slot
		go func(url string) {
			defer func() {
				<-running // releases slot
			}()
			// do work
		}(url)
	}
```

<span class="fragment current-only" data-code-focus="4-8"></span>
