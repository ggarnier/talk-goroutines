## Deadlock

```shell
$ go run main.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        main.go:12 +0xaa

goroutine 17 [chan receive]:
main.main.func1(0xc4200720c0, 0xc420072060)
        main.go:8 +0x3e
created by main.main
        main.go:7 +0x89
exit status 2
```

<span class="fragment current-only" data-code-focus="2"></span>
