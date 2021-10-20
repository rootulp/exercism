package account

type Account struct {
}

func Open(initialDeposit int64) *Account {
	return &Account{}
}

func (*Account) Close() (payout int64, ok bool) {
	return 0, true
}

func (*Account) Balance() (balance int64, ok bool) {
	return 0, true
}

func (*Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return 0, true
}
