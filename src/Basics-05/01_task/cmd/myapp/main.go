package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust)

	///
	price, err := internal.CalcPrice(cust, 100)
	fmt.Printf("%d %v\n", price, err)

	///
	cust.Discount = false
	price, err = internal.CalcPrice(cust, 1000)
	fmt.Printf("%d %v\n", price, err)

	///
	cust.Discount, cust.Debt = true, 0
	price, err = internal.CalcPrice(cust, 10000)
	fmt.Printf("%d %v\n", price, err)
}
