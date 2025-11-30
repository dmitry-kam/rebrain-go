package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

func main() {
	cust := internal.NewCustomer("Johnny", 23, 10000, 1000, true)
	_ = startTransaction(cust)
	fmt.Printf("%+v\n", cust)

	partner := internal.NewPartner("Donald", 23, 10000, 1000)
	_ = startTransaction(partner)
	fmt.Printf("%+v\n", partner)

	err := startTransactionDynamic(partner)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", partner)
	}

	err = startTransactionDynamic(123)
	if err != nil {
		fmt.Println(err)
	}
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}

func startTransactionDynamic(d interface{}) error {
	debtor, ok := d.(internal.Debtor)

	if ok {
		return startTransaction(debtor)
	}

	return errors.New("incorrect type")
}
