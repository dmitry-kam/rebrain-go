package main

import (
	"fmt"
	v1 "utils"
	v3 "utils/v3"
)

func main() {
	fruits := []string{"apple", "banana", "orange"}
	numbers := []int{10, 20, 30, 40}

	fmt.Printf("Using v1.1.0:\n")
	fmt.Printf("Contains 'banana': %t\n", v1.Contains(fruits, "banana"))
	fmt.Printf("Contains 'grape': %t\n", v1.Contains(fruits, "grape"))
	fmt.Printf("ContainsInt 30: %t\n", v1.ContainsInt(numbers, 30))
	fmt.Printf("ContainsInt 99: %t\n", v1.ContainsInt(numbers, 99))

	fmt.Printf("Using v3.0.0:\n")
	fmt.Printf("InSlice 'banana': %t\n", v3.InSlice(fruits, "banana"))
	fmt.Printf("InSlice 'grape': %t\n", v3.InSlice(fruits, "grape"))
	fmt.Printf("InSliceInt 30: %t\n", v3.InSliceInt(numbers, 30))
	fmt.Printf("InSliceInt 99: %t\n", v3.InSliceInt(numbers, 99))
}
