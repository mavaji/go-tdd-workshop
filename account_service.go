package go_tdd_workshop

type AccountService struct {
	balance int
}

func (s *AccountService) Deposit(amount int) {
	s.balance += amount
}

func (s *AccountService) getBalance() int {
	return s.balance
}

func (s *AccountService) Withdraw(amount int) {
	s.balance -= amount
}

func NewAccountService() *AccountService {
	return &AccountService{}
}
