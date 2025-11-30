package main

import (
	"fmt"
	"myapp/internal"
)

type Printer interface {
	print()
}

type A struct {
}

func (a *A) print() {
	fmt.Println("Hello, playground")
}

func main() {
	cust := internal.NewCustomer("Johnny", 23, 10000, 1000, true)
	_ = startTransaction(cust)
	fmt.Printf("%+v\n", cust)

	partner := internal.NewPartner("Donald", 23, 10000, 1000)
	_ = startTransaction(partner)
	fmt.Printf("%+v\n", partner)

	///////////////////
	//instance := A{}
	//
	//var p Printer
	//// cannot use instance (variable of type A) as Printer value in assignment: A does not implement Printer (method print has pointer receiver)
	//p = instance
	//p.print()
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
