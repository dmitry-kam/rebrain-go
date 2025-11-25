package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	var cust internal.Customer
	//cust := new(internal.Customer)
	//cust := internal.Customer{}

	fmt.Printf("%+v\n", cust) // {Name: Age:0 Balance:0 Debt:0 Discount:false CalcDiscount:<nil>}

	cust1 := internal.Customer{
		Age:     23,
		Balance: 10000,
		Debt:    1000,
		Name:    "Donald",
	}

	fmt.Printf("%+v\n", cust1)      //{Name:Donald Age:23 Balance:10000 Debt:1000 Discount:false CalcDiscount:<nil>}
	fmt.Printf("%+v\n", cust1.Name) // Donald

	cust1.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	discount, _ := cust1.CalcDiscount()

	fmt.Printf("%d\n", discount) // 0

	////////////////////////////////////
	cust2 := internal.NewCustomer("Dmitrii", 34, 100000, 0, true)
	// constructor returns not a copy of the instance, but a pointer to it. This is also possible—in this situation,
	// it makes no sense to create a structure and then copy it as a return value (an unnecessary memory allocation),
	// so we can simply return the pointer. This can potentially improve performance on large structures.

	cust2.CalcDiscount = func() (int, error) {
		if !cust2.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust2.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Printf("%+v\n", cust2)

	cust3 := cust2
	cust3.Age = 88
	fmt.Printf("%+v\n", cust2)
	fmt.Printf("%+v\n", cust3)
	//&{Name:Dmitrii Age:88 Balance:100000 Debt:0 Discount:true CalcDiscount:0x482620}
	//&{Name:Dmitrii Age:88 Balance:100000 Debt:0 Discount:true CalcDiscount:0x482620}

	// not pointers, new structures
	cust4 := internal.AbsolutelyNewCustomer("Dmitrii", 34, 100000, 0, true)
	cust5 := cust4
	cust5.Age = 100
	fmt.Printf("%+v\n", cust4)
	fmt.Printf("%+v\n", cust5)
	//{Name:Dmitrii Age:34 Balance:100000 Debt:0 Discount:true CalcDiscount:<nil>}
	//{Name:Dmitrii Age:100 Balance:100000 Debt:0 Discount:true CalcDiscount:<nil>}

	/////////////////////// immutable

	config := internal.NewConfig()
	fmt.Printf("%+v\n", config)
	//config.APIKey = "new"
	// config.private = 2// restricted
	fmt.Printf("%+v\n", config)
	fmt.Printf("%s\n", config.APIKey())

}
