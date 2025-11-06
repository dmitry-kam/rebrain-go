package main

import (
	"fmt"
	"math"
)

func main() {
	var pi rune = 'π'
	fmt.Printf("%c = %.48f\n", pi, math.Pi)

	R := new(float64)
	L := 35.0
	*R = L / (2 * math.Pi)
	fmt.Printf("Circum: circumference = %.2f\n", L)
	fmt.Printf("R = %.2f\n", *R)
	fmt.Printf("S = %.3f\n", math.Round(math.Pi*math.Pow(*R, 2)*100)/100)
}
