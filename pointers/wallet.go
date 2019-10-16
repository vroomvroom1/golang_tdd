package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is custom type that holds int
type Bitcoin int

// Stringer is interface that specifies string formatting
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a struct to hold the values
type Wallet struct {
	balance Bitcoin
}

// Deposit accepts an amount and adds it to our wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance retrieves the wallet amount
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds is repeatable error message
var ErrInsufficientFunds = errors.New("cannon withdraw, insufficient funds")

// Withdraw allows you to withdraw from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
