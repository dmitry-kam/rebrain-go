package main

import (
	"fmt"
)

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("Christian week %v\n", days)
	weekDays := make([]string, 5)

	copy(weekDays, days[1:6])
	days = append(days[6:], days[:1]...)

	fmt.Printf("Working days %v\n", weekDays)
	fmt.Printf("Weekend: %v\n", days)

	sovietWeek := append(weekDays, days...)
	fmt.Printf("Soviet week %v\n", sovietWeek)
}
