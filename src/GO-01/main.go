package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Printf("Hello. Current time is %02d-%02d-%d %02d:%02d\n",
		currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
}
