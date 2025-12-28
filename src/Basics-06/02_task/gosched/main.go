package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Start")
	go func() {
		fmt.Println("Start goroutine")
		for {
			someWorkload()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}

func someWorkload() {
	//before
	//for {
	//}
	//after
	for {
		// Any case creates cooperation points and terminates the program:
		runtime.Gosched()
		//time.Sleep(0)
		//_ = time.Now().UnixNano() // System call
		//_ = make(chan int) // Memory allocation
		//var m sync.Mutex
		//m.Lock()
		//m.Unlock() // Synchronization primitive
	}
}
