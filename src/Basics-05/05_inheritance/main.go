package main

import (
	"fmt"
	"time"
)

type Card struct {
	Balance     int
	ExpiredDate string
	CVV         int
	Num         int64
	Owner       string
}

type CreditCard struct {
	Card
	Limit int
}

type Checker interface {
	CheckDate() bool
}

func (c *Card) CheckDate() bool {
	ex, _ := time.Parse("2006.01.02", c.ExpiredDate)
	return time.Now().Before(ex)
}

type CreditCard1 struct {
	*Card
	Limit int
}

type CreditCard2 struct {
	Checker
	Limit int
}

type Buyer interface {
	buy()
}

type Seller interface {
	sell()
}

type Dealer interface {
	Buyer
	Seller
}

func main() {
	cc := &CreditCard{
		Card{
			Balance:     10000,
			ExpiredDate: "01.02.2023",
			CVV:         132,
			Num:         4234536475474656,
			Owner:       "Vasily Ivanov",
		},
		100000,
	}

	cc1 := &CreditCard1{
		&Card{
			Balance:     10000,
			ExpiredDate: "01.02.2023",
			CVV:         132,
			Num:         4234536475474656,
			Owner:       "Vasily Ivanov",
		},
		100000,
	}

	fmt.Printf("%+v\n", cc)
	fmt.Printf("%+v\n", cc1)

	fmt.Printf("Limit: %d, Owner: %s\n", cc.Limit, cc.Owner)

	cc.Owner = "Oleg"
	cc.Limit = 0

	fmt.Printf("Limit: %d, Owner: %s\n\n", cc.Limit, cc.Owner)

	///////////
	c1 := &Card{
		Balance:     10000,
		ExpiredDate: "01.02.2023",
		CVV:         132,
		Num:         4234536475474656,
		Owner:       "Vasily Ivanov",
	}
	cc3 := &CreditCard2{
		c1,
		100000,
	}

	fmt.Printf("%t", cc3.CheckDate())

	//cc3.Owner = "Dima"
	// cc3.Owner undefined (type *CreditCard has no field or method Owner)
}
