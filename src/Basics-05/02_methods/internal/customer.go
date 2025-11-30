package internal

import (
	"errors"
)

type Customer struct {
	Name         string
	Age          int
	balance      int
	debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func (c Customer) WrOffDebt() (Customer, error) {
	if c.debt >= c.balance {
		return c, errors.New("Not possible write off")
	}

	c.balance -= c.debt
	c.debt = 0

	return c, nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) Customer {
	return Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		Discount: discount,
	}
}

func (c *Customer) WrOffDebtPointer() error {
	if c.debt >= c.balance {
		return errors.New("Not possible write off")
	}

	c.balance -= c.debt
	c.debt = 0

	return nil
}

func NewCustomerPointer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		Discount: discount,
	}
}
