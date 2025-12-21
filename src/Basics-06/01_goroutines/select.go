package main

import (
	"fmt"
	"time"
)

func fibonacci2channels(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		// A select blocks until one of its cases can run, then it executes that case.
		// It chooses one at random if multiple are ready.
		case <-quit:
			fmt.Println("quit")
			return
		// The default case in a select is run if no other case is ready.
		default:
			elapsed := func() time.Duration {
				return time.Since(time.Now()).Round(time.Millisecond)
			}
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2channels(c, quit)
}
