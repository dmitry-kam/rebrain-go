package main

import "fmt"

func main() {
	ch1 := make(chan float64)       // unbuffered (float64)
	fmt.Println(len(ch1), cap(ch1)) // 0, 0
	//ch1 <- 3.14
	//fmt.Println(len(ch1), cap(ch1)) // fatal error: all goroutines are asleep - deadlock!
	// This happens when a goroutine is blocked and can't release
	// the lock. In this case, it happened because we had one goroutine
	// that encountered an unbuffered channel and couldn't proceed.

	go func() {
		val, ok := <-ch1
		fmt.Println(val, ok)
	}()
	ch1 <- 3.14

	//////////////

	ch2 := make(chan int, 3)        // buffered, capacity 3 (int)
	fmt.Println(len(ch2), cap(ch2)) // 0, 3
	ch2 <- 4
	fmt.Println(len(ch2), cap(ch2)) // 1, 3

	close(ch1)
	result, ok := <-ch1
	fmt.Println(result, ok) // 0, false

	result1, ok := <-ch2
	fmt.Println(result1, ok) // 4, true
}
