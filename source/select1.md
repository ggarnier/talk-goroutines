## Multiplexing channels

```go
	select {
	case v := <-ch1:
		fmt.Println("value read from ch1")
	case v := <-ch2:
		fmt.Println("value read from ch2")
	case ch1 <- 10:
		fmt.Println("value sent to ch1")
	default:
		fmt.Println("channels not ready")
	}
```

If more than one condition is ready, a `case` is randomly chosen

<span class="fragment current-only" data-code-focus="2,4,6"></span>
<span class="fragment current-only" data-code-focus="8"></span>
