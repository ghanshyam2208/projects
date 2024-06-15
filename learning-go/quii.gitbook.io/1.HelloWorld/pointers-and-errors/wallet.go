package pointersanderrors

import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func NewWallet() Wallet {
	var wallet Wallet = Wallet{
		balance: 0,
	}
	return wallet
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(deductAmount Bitcoin) error {
	if w.balance < deductAmount {
		return ErrInsufficientFunds
	}
	w.balance = w.balance - deductAmount
	return nil
}
