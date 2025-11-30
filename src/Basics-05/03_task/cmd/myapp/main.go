package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust1 := internal.NewCustomer("Siegfried", 23, 10000, 1000, true)
	cust2 := internal.NewCustomer("Njáll", 23, 10000, 1000, false)
	cust3 := internal.NewCustomer("Egil", 23, 10000, 0, true)

	///
	price, err := internal.CalcPrice(cust1, 100)
	fmt.Printf("%d %v\n", price, err)

	///
	price, err = internal.CalcPrice(cust2, 1000)
	fmt.Printf("%d %v\n", price, err)

	///
	price, err = internal.CalcPrice(cust3, 10000)
	fmt.Printf("%d %v\n", price, err)
}
