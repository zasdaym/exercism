// Package account provides Bank Account solution.
package account

import "sync"

// Account represents a bank account.
type Account struct {
	mu      sync.Mutex
	balance int64
	closed  bool
}

// Open opens a new Account.
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}

	return &Account{
		balance: deposit,
	}
}

// Close closes a bank account.
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

// Deposit deposits amount to the Account's balance.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	updatedBalance := a.balance + amount
	if updatedBalance < 0 {
		return 0, false
	}

	a.balance = updatedBalance
	return a.balance, true
}

// Balance returns the balance of a bank account.
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}
