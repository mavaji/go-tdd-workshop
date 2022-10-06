package go_tdd_workshop

type AccountService struct {
}

func (s *AccountService) Deposit(amount int) {

}

func (s *AccountService) getBalance() int {
	return 5
}

func NewAccountService() *AccountService {
	return &AccountService{}
}
