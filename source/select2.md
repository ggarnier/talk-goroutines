## Unblocking receives

```go
	ch := make(chan int)
	select {
	case v, ok := <-ch:
		if ok {
			fmt.Printf("Value is %d", v)
		} else {
			fmt.Print("channel is closed")
		}
	default:
		fmt.Print("channel is empty")
	}
```
