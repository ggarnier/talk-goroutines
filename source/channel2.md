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
