package main

import (
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Johnny", 23, 10000, 1000, true)
	_ = startTransaction(cust)
	fmt.Printf("%+v %d %d\n", cust, cust.GetBalance(), cust.GetDebt())
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
