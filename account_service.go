package go_tdd_workshop

import "errors"

type AccountService struct {
	balance int
}

func (s *AccountService) Deposit(amount int) {
	s.balance += amount
}

func (s *AccountService) getBalance() int {
	return s.balance
}

func (s *AccountService) Withdraw(amount int) error {
	if s.balance >= amount {
		s.balance -= amount
		return nil
	}

	return errors.New("insufficient funds")
}

func NewAccountService() *AccountService {
	return &AccountService{}
}
