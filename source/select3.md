## Timing out

```go
	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Printf("Value is %d", v)
	case <- time.After(2*time.Second):
		fmt.Print("no data after 2 seconds")
	}
```

<span class="fragment current-only" data-code-focus="5"></span>
