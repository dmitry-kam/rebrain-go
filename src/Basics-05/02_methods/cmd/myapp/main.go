package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	cust, _ = cust.WrOffDebt()
	fmt.Printf("%+v\n", cust)

	custP := internal.NewCustomerPointer("Dmitry", 23, 10000, 1000, true)
	custP.WrOffDebtPointer()
	fmt.Printf("%+v\n", custP)
}
