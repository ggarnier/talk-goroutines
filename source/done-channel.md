### Forcing goroutines to stop

Using a `done` channel

<!-- examples/done-channel/main.go -->

```go
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		bufio.NewReader(os.Stdin).ReadByte() // read input from stdin
	}()

	randomNumbers := func() <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for {
				select {
				case <-done:
					return
				default:
					stream <- rand.Intn(100)
				}
			}
		}()
		return stream
	}
```

<span class="fragment current-only" data-code-focus="1,4"></span>
<span class="fragment current-only" data-code-focus="14-19"></span>
