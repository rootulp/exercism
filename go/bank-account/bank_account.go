package account

import "sync"

type Account struct {
	balance  int64
	isClosed bool
	sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isClosed {
		return 0, false
	}
	a.isClosed = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.isClosed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isClosed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
