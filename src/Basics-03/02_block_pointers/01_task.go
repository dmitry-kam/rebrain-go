package main

import (
	"fmt"
)

func main() {
	A := new(int)
	B := 100

	A = &B
	fmt.Println(*A)

	*A = 200
	fmt.Println(B)
}
