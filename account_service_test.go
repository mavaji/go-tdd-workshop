package go_tdd_workshop

import "testing"

//- [ ] can't have negative balance
//- [ ] can't have more than a maximum balance
//- [ ] can handle different customers (account ids)
//- [ ] can work with a database
//- [ ] can't have empty or invalid account id
//- [ ] can't have more than a maximum deposit
//- [ ] can't have more than a maximum withdraw

func TestAccountService(t *testing.T) {
	t.Run("can deposit money into account", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Deposit(5)

		want := 5
		got := accountService.getBalance()
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("can withdraw money from account", func(t *testing.T) {
		accountService := NewAccountService()
		accountService.Deposit(5)
		accountService.Withdraw(4)

		want := 1
		got := accountService.getBalance()
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}
