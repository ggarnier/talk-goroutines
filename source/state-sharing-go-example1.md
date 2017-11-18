## Example: Healthchecks in Go

```go
	type Results struct {
		okCount  int
		errCount int
	}

	func main() {
		urls := []string{"url1", "url2", ...}

		r := runChecks(urls)
		fmt.Printf("%d ok, %d errors\n", r.okCount, r.errCount)
	}
```

<span class="fragment current-only" data-code-focus="7,9"></span>
