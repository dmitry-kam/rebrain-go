package internal

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	// the same
	//pointer := Customer{
	//	Name:     name,
	//	Age:      age,
	//	Balance:  balance,
	//	Debt:     debt,
	//	Discount: discount,
	//}
	//return &pointer

	// short
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func AbsolutelyNewCustomer(name string, age int, balance int, debt int, discount bool) Customer {
	return Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}
