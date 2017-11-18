## Race Detector tool

```shell
	$ go build -race
	$ ./main
	==================
	WARNING: DATA RACE
	Read at 0x00c4200ee2c0 by main goroutine:
		main.runChecks()
				main.go:35 +0x16b
		main.main()
				main.go:45 +0x87

	Previous write at 0x00c4200ee2c0 by goroutine 7:
		main.runChecks.func1()
				main.go:26 +0x178

	Goroutine 7 (finished) created at:
		main.runChecks()
				main.go:18 +0x123
		main.main()
				main.go:45 +0x87
	==================
	==================
	WARNING: DATA RACE
	...
	==================
	Found 2 data race(s)
```

<span class="fragment current-only" data-code-focus="1-2"></span>
