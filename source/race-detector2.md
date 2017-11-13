## Race Detector tool

```shell
	$ go build -race
	$ ./main
	==================
	WARNING: DATA RACE
	Read at 0x00c4200ee2c0 by main goroutine:
		main.runChecks()
				/Users/ggarnier/dev/pessoal/talk-goroutines/x.go:35 +0x16b
		main.main()
				/Users/ggarnier/dev/pessoal/talk-goroutines/x.go:45 +0x87

	Previous write at 0x00c4200ee2c0 by goroutine 7:
		main.runChecks.func1()
				/Users/ggarnier/dev/pessoal/talk-goroutines/x.go:26 +0x178

	Goroutine 7 (finished) created at:
		main.runChecks()
				/Users/ggarnier/dev/pessoal/talk-goroutines/x.go:18 +0x123
		main.main()
				/Users/ggarnier/dev/pessoal/talk-goroutines/x.go:45 +0x87
	==================
	==================
	WARNING: DATA RACE
	...
	==================
	Found 2 data race(s)
```

<span class="fragment current-only" data-code-focus="1-2"></span>
