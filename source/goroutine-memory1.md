## Goroutines memory usage

```go
	memConsumed := func() uint64 {...}
	var ch <-chan interface{}
	var wg sync.WaitGroup

	const numGoroutines = 100000
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := 0; i < numGoroutines; i++ {
		go func() { wg.Done(); <-ch }()
	}
	wg.Wait()
	after := memConsumed()

	fmt.Printf("%.3f bytes", float64(after-before)/numGoroutines)
```

<span class="fragment current-only" data-code-focus="5,9"></span>
<span class="fragment current-only" data-code-focus="10"></span>
<span class="fragment current-only" data-code-focus="8,13"></span>

<p style="font-size: 70%; margin-top: 50px">source: Concurrency in Go, page 44</p>
