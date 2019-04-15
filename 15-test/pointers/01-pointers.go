package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

func (wallet *Wallet) Deposit(amount Bitcoin) { //wallet is a copy of whatever we called the method from
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) (error, Bitcoin) {
	if amount > wallet.balance {
		return errors.New("no enough funds"), wallet.balance
	}
	return nil, wallet.balance - amount
}
