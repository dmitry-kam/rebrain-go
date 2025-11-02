package main

import (
	"fmt"
	"strconv"
)

var one = "104"
var two = 35

func main() {
	convertedOne, _ := strconv.Atoi(one)
	convertedTwo := strconv.Itoa(two)
	fmt.Printf("%T %s as %T: %d\n", one, one, convertedOne, convertedOne)
	fmt.Printf("%T %d as %T: %q\n", two, two, convertedTwo, convertedTwo)
}
