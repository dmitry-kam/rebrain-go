package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexBankCell struct {
	DollarCell float64
	mu         sync.Mutex
}

func (b *MutexBankCell) GetBalance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.DollarCell
}

func (b *MutexBankCell) SubBalance(value float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.DollarCell-value >= 0 {
		time.Sleep(time.Millisecond * 1)
		b.DollarCell -= value
	}
}

func (b *MutexBankCell) AddBalance(value float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.DollarCell += value
}

func main() {
	bc := &MutexBankCell{
		DollarCell: 10.0,
	}

	go bc.SubBalance(7.0)
	bc.SubBalance(7.0)
	time.Sleep(time.Second)

	fmt.Println(bc.GetBalance())
}
