package pointers

import (
	"errors"
	"fmt"
)

type Monero int

type Wallet struct {
	balance Monero
}

func (xmr Monero) String() string {
	return fmt.Sprintf("%d XMR", xmr)
}

func (wallet *Wallet) Deposit(amount Monero) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Monero {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Monero) error {

	if amount > wallet.balance {
		return errors.New("can not withdraw, insufficient funds")
	}

	wallet.balance -= amount

	return nil
}
