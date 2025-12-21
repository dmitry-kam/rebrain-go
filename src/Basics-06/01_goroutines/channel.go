package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered
	ch <- 1                 // write to channel
	ch <- 2
	//ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch) // get from channel
	fmt.Println(<-ch)
	ch <- 3
	printFromChannel(ch, false)
	ch <- 4
	printFromChannel(ch, true)
	// ch <- 5 // panic: send on closed channel
	printFromChannel(ch, false) // v 0 is undefined. channel has been closed

	fmt.Println("-------------------")

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	printAllValuesFromChannel(c)

	fmt.Println("-------------------")
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Note: Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are
	// no more values coming, such as to terminate a range loop.
	close(c)
}

func printFromChannel(c chan int, thenClose bool) {
	v, ok := <-c
	fmt.Printf("printFromChannel %t\n", ok)

	if !ok {
		fmt.Printf("v %d is undefined. channel has been closed\n", v)
	} else {
		fmt.Println(v)
		if thenClose {
			close(c)
		}
	}
}

func printAllValuesFromChannel(c chan int) {
	for i := range c {
		fmt.Printf("printFromChannel %d\n", i)
	}
}
