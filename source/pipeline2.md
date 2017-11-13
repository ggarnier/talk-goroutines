## Pipelines

<!-- examples/pipeline/main.go -->

```go
	// Inputs an infinite stream
	randomNumbers := func() <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for {
				stream <- rand.Intn(100)
			}
		}()
		return stream
	}

	// Transforms
	double := func(input <-chan int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range input {
				stream <- i * 2
			}
		}()
		return stream
	}
```

<span class="fragment current-only" data-code-focus="2"></span>
<span class="fragment current-only" data-code-focus="14"></span>
<span class="fragment current-only" data-code-focus="3,10,15,22"></span>
<span class="fragment current-only" data-code-focus="5,17"></span>
<span class="fragment current-only" data-code-focus="6-8,18-20"></span>
