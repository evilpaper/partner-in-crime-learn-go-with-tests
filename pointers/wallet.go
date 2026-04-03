package pointers

type Wallet struct {
	balance int
}

// Note the star (*), it means "a pointer to a Wallet"
func (wallet *Wallet) Deposit(amount int) {
	// Note that we don't need to dereference the pointer, it works automatically. Just how the language works.
	wallet.balance += amount
}

func (wallet *Wallet) Balance() int {
	return wallet.balance
}