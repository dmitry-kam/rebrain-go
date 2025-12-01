package internal

import (
	"errors"
)

type Customer struct {
	*Overdue
	Name         string
	Age          int
	Discount     bool
	CalcDiscount func() (int, error)
}

func (c *Customer) WrOffDebt() error {
	if c.Overdue.debt >= c.Overdue.balance {
		return errors.New("not possible write off")
	}

	c.Overdue.balance -= c.Overdue.debt
	c.Overdue.debt = 0

	return nil
}

func (c *Customer) GetDebt() int {
	return c.Overdue.debt
}

func (c *Customer) GetBalance() int {
	return c.Overdue.balance
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Discount: discount,
		Overdue:  NewOverdue(balance, debt),
	}
}
