## Closed channels

How to receive until the channel closes?

```go
	ch := make(chan string, 2)

	go func() {
		ch <- "foo"
		ch <- "bar"
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
```

If the channel is empty, `range` blocks until the channel is closed

<span class="fragment current-only" data-code-focus="9"></span>
