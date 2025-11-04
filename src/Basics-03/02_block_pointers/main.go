package main

import (
	"fmt"
)

func main() {
	var x int

	pntr := &x
	fmt.Println(pntr)

	x = 64
	fmt.Println(*pntr)

	*pntr = 31
	fmt.Println(*pntr)
	fmt.Println(x)

	fmt.Printf("x: type=%T value=%d\n", x, x)
	fmt.Printf("pntr: type=%T value=%d address=%p\n", pntr, *pntr, pntr)

	var intValue int
	pntr = &intValue
	fmt.Println(pntr)

	pntr2 := new(int)
	fmt.Println(pntr2)
}
