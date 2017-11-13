package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	done := make(chan struct{})
	go func() {
		defer func() {
			fmt.Println("finishing...")
			done <- struct{}{}
		}()
		bufio.NewReader(os.Stdin).ReadByte()
	}()

	randomNumbers := func() <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for {
				select {
				case <-done:
					return
				default:
					stream <- rand.Intn(100)
				}
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
