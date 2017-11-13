package main

import (
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var sleep bool
	if len(os.Args) > 1 && os.Args[1] == "sleep" {
		sleep = true
	}
	numGoroutines := 2
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				for j := 0; j < 10000000; j++ {
				}
				if sleep {
					time.Sleep(1 * time.Nanosecond)
				}
			}
		}()
	}

	wg.Wait()
}
