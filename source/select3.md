## Timing out

```go
	ch := make(chan int)
	select {
	case v := <-ch:
	case <- time.After(2*time.Second):
		fmt.Print("no data after 2 seconds")
	}
```
