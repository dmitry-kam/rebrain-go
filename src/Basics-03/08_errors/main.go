package main

import "fmt"

func main() {
	_, err := checkNum(11)
	if err != nil {
		fmt.Println("invalid num:", err)
	}
}

func checkNum(x int) (valid bool, err error) {
	if x > 10 {
		return false, fmt.Errorf("out of positive range: %v", x)
	}
	if x < -10 {
		return false, fmt.Errorf("out of negative range: %v", x)
	}

	return true, nil
}
