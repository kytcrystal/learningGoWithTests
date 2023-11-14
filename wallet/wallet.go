package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
	//private balance
	//so only this package can manipulate the value
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += Bitcoin(amount)
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if wallet.balance < Bitcoin(amount) {
		return ErrInsufficientFunds
	}
	wallet.balance -= Bitcoin(amount)
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance

}
