## Buffered channels

- created with a maximum capacity
- sender and receiver block each other when it's full
- no guarantee of delivery

```go
	stream := make(chan int, 1)
	go func() {
		stream <- 1
		stream <- 2 // blocks until the first receiver
	}()
	go func() {
		<-stream // blocks until the first sender
		<-stream
	}()
```
