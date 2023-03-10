package banking

// Account in Banking
type account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 123}
	return &account
}

func (tAccount *account) Deposit(cash int) {
	tAccount.balance = cash
}

func (tAccount account) GetDeposit() int {
	return tAccount.balance
}

func (tAccount *account) AddDeposit(cash int) {
	tAccount.balance += cash
}
