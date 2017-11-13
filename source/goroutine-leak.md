## Goroutine leak

```go
	urls := make(chan string)

	go func() {
		// defer close(urls)
		urls <- "http://example.com/1"
		urls <- "http://example.com/2"
	}()

	go func() {
		defer fmt.Println("This will never run")
		for url := range urls {
			http.Get(url)
			fmt.Println(url)
		}
	}()
```

<span class="fragment current-only" data-code-focus="4"></span>
<span class="fragment current-only" data-code-focus="10"></span>
