package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Note the star (*), it means "a pointer to a Wallet"
func (wallet *Wallet) Deposit(amount Bitcoin) {
	// Note that we don't need to dereference the pointer, it works automatically. Just how the language works.
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}