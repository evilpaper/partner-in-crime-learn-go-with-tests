package pointers

import (
	"fmt"
)

type Wallet struct {
	balance int
}

// Notice the start (*). Read it as "a pointer to a Wallet"
func (wallet *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &wallet.balance)
	// We don't need to dereference the pointer here because we are working with a struct. Just how the language works.
	wallet.balance += amount
}

func (wallet *Wallet) Balance() int {
	return wallet.balance
}