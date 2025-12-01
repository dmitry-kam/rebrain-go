package internal

type Overdue struct {
	balance int
	debt    int
}

func NewOverdue(balance int, debt int) *Overdue {
	return &Overdue{
		balance: balance,
		debt:    debt,
	}
}
