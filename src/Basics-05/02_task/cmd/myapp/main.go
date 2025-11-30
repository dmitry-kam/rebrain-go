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
	price, err := cust1.CalcPrice(100)
	fmt.Printf("%d %v\n", price, err)

	///
	price, err = cust2.CalcPrice(1000)
	fmt.Printf("%d %v\n", price, err)

	///
	price, err = cust3.CalcPrice(10000)
	fmt.Printf("%d %v\n", price, err)
}
