package go_tdd_workshop

import "errors"

type AccountService struct {
	db DB
}

func (s *AccountService) Deposit(accountId string, amount int) error {
	if s.db.GetBalanceByAccountId(accountId)+amount > 100000 {
		return errors.New("too rich")
	}

	s.db.UpdateBalanceByAccountId(accountId, amount)
	return nil
}

func (s *AccountService) getBalance(accountId string) int {
	return s.db.GetBalanceByAccountId(accountId)
}

func (s *AccountService) Withdraw(accountId string, amount int) error {
	if s.db.GetBalanceByAccountId(accountId) >= amount {
		s.db.UpdateBalanceByAccountId(accountId, -amount)
		return nil
	}

	return errors.New("insufficient funds")
}

func NewAccountService(db DB) *AccountService {
	return &AccountService{db}
}

type DB interface {
	GetBalanceByAccountId(accountId string) int
	UpdateBalanceByAccountId(accountId string, amount int)
}
