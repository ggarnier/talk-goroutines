### Version 1: sequential code

```go
	func runChecks(urls []string) Results {
		r := Results{}
		for _, url := range urls {
			resp, err := http.Head(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				r.errCount++
			} else {
				r.okCount++
			}
		}

		return r
	}
```

<!--
- gives the correct response
- no sync mecanisms required
- hard to scale
-->
