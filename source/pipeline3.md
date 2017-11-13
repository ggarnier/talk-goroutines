## Pipelines

```go
	// Filters
	onlyMultiplesOf10 := func(input <-chan int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range input {
				if i%10 == 0 {
					stream <- i
				}
			}
		}()
		return stream
	}

	pipeline := onlyMultiplesOf10(double(randomNumbers()))
	for i := range pipeline { // Infinite loop
		fmt.Println(i)
	}
```

<span class="fragment current-only" data-code-focus="15-18"></span>
