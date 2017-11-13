package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randomNumbers := func() <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := 1; i <= 30; i++ {
				stream <- rand.Intn(100)
			}
		}()
		return stream
	}

	double := func(input <-chan int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range input {
				stream <- i * 2
			}
		}()
		return stream
	}

	onlyMultiplesOf10 := func(input <-chan int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range input {
				if i%10 == 0 {
					stream <- i
				}
			}
		}()
		return stream
	}

	pipeline := onlyMultiplesOf10(double(randomNumbers()))
	for i := range pipeline {
		fmt.Println(i)
	}
}
