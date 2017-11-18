## Unbuffered channels

- sender and receiver block each other
- provide guarantee of delivery

```go
	stream := make(chan int)
	go func() {
		stream <- 1
	}()
	go func() {
		<-stream
	}()
```

<span class="fragment current-only" data-code-focus="1"></span>
<span class="fragment current-only" data-code-focus="3,6"></span>
