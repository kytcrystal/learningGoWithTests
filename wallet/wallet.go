package wallet

import "fmt"

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

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance

}
