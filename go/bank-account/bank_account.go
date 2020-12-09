// Package account is solution for problem Bank Account.
package account

import (
	"sync"
)

// Account represents a bank account.
type Account struct {
	sync.Mutex
	balance int64
	closed  bool
}

// Open creates new Account.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	account := &Account{balance: initialDeposit}
	return account
}

// Close closes an Account.
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

// Balance returns balance of an Account.
func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit deposits an Account's balance.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
