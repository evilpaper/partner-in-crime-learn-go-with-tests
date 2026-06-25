package pointers

import (
	"errors"
	"fmt"
)

// We want our wallet to contain Bitcoins. We could use just raw int but we want our wallet to know the concept of a Bitcoin.
// So we create a new type called Bitcoin that is an int.
// int is good for counting things.
// Go lets us create new types from existing types. So here we have created a new type called Bitcoin that is an int.
type Bitcoin int

type Wallet struct {
	// Note the lowercase balance. In Go, lowercase means it's private. That is not exported to other packages.
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Note the star (*), it means "a pointer to a Wallet"
func (wallet *Wallet) Deposit(amount Bitcoin) {
	// Note that we don't need to dereference the pointer, it works automatically. Just how the language works.
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return errors.New("oh no")
	}
   	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
