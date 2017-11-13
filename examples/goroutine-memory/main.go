package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

func main() {
	var ch <-chan interface{}
	var wg sync.WaitGroup

	const numGoroutines = 100000
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := 0; i < numGoroutines; i++ {
		go func() { wg.Done(); <-ch }()
	}
	wg.Wait()

	after := memConsumed()
	fmt.Printf("%d goroutines - %.3f bytes\n", numGoroutines, float64(after-before)/numGoroutines)
}
