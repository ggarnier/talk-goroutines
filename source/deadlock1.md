## Deadlock

```go
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		n := <-ch2
		ch1 <- 2 * n
	}()

	n := <-ch1
	ch2 <- 2 * n
```

<span class="fragment current-only" data-code-focus="5,9"></span>
