package main

import (
	"fmt"
)

func main() {
	weekDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("Christian week %v\n", weekDays)
	workingDays := make([]string, 5)

	copy(workingDays, weekDays[1:6])
	weekDays = append(weekDays[6:], weekDays[:1]...)

	fmt.Printf("Working days %v\n", workingDays)
	fmt.Printf("Weekend: %v\n", weekDays)

	sovietWeek := append(workingDays, weekDays...)
	fmt.Printf("Soviet week %v\n", sovietWeek)
}
