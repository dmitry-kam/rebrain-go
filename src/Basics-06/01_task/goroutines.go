package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n1, n2 := 44, 45

	startTime := time.Now()
	fibN1, fibN2 := fib(n1), fib(n2)
	duration := time.Since(startTime)
	fmt.Printf("\rFibonacci(%d) = %d\n", n1, fibN1)
	fmt.Printf("\rFibonacci(%d) = %d\n", n2, fibN2)
	fmt.Printf("\rCounted for %f\n", duration.Seconds()) // Synchronously: 10.435821

	startTime = time.Now()
	ch1, ch2 := make(chan int), make(chan int)
	go asyncFib(n1, ch1)
	go asyncFib(n2, ch2)
	fibN1, fibN2 = <-ch1, <-ch2
	duration = time.Since(startTime)
	fmt.Printf("\rFibonacci(%d) = %d\n", n1, fibN1)
	fmt.Printf("\rFibonacci(%d) = %d\n", n2, fibN2)
	fmt.Printf("\rCounted for %f\n", duration.Seconds()) // Asynchronously: 6.560279
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func asyncFib(x int, c chan int) {
	c <- fib(x)
	close(c)
}
