### Forcing goroutines to stop

Using context

```go
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		defer cancelFunc()
		bufio.NewReader(os.Stdin).ReadByte() // read input from stdin
	}()

	randomNumbers := func() <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					stream <- rand.Intn(100)
				}
			}
		}()
		return stream
	}
```

<span class="fragment current-only" data-code-focus="1"></span>
<span class="fragment current-only" data-code-focus="3,13"></span>
