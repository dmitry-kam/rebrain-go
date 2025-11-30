package internal

import "errors"

const defaultDiscount = 500

type Customer struct {
	Name     string
	Age      int
	Balance  int
	Debt     int
	discount bool
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

func (cust *Customer) calcDiscount() (int, error) {
	if !cust.discount {
		return 0, errors.New("discount not available")
	}
	result := defaultDiscount - cust.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func (cust *Customer) CalcPrice(price int) (int, error) {
	discount, err := cust.calcDiscount()
	if err != nil {
		return 0, err
	}

	if price > discount {
		return price - discount, nil
	}

	return 0, nil
}
