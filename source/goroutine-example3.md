```go
	numGoroutines := 2
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				for j := 0; j < 10000000; j++ {
				}
				time.Sleep(1 * time.Nanosecond)
			}
		}()
	}

	wg.Wait()
```

<span class="fragment current-only" data-code-focus="11"></span>
