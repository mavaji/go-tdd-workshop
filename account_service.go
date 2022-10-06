package go_tdd_workshop

import "errors"

type AccountService struct {
	balance map[string]int
}

func (s *AccountService) Deposit(accountId string, amount int) error {
	if s.balance[accountId]+amount > 100000 {
		return errors.New("too rich")
	}

	s.balance[accountId] += amount
	return nil
}

func (s *AccountService) getBalance(accountId string) int {
	return s.balance[accountId]
}

func (s *AccountService) Withdraw(accountId string, amount int) error {
	if s.balance[accountId] >= amount {
		s.balance[accountId] -= amount
		return nil
	}

	return errors.New("insufficient funds")
}

func NewAccountService() *AccountService {
	return &AccountService{make(map[string]int)}
}
