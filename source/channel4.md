## Closed channels

- We can't send values anymore
- We can still receive values sent before closing it
- If empty, receiving values gives the zero value of the channel type

```go
	ch := make(chan string, 2)
	ch <- "foo"
	ch <- "bar"
	close(ch)

	fmt.Println(<-ch) // "foo"
	fmt.Println(<-ch) // "bar"
	fmt.Println(<-ch) // ""
```
