package repositories

import (
	"banking_app2/cmd/internals/dto"
	"time"
)

type Account struct {
	Id        int64     `db:"id"`
	Owner     string    `db:"owner"`
	Balance   int64     `db:"balance"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

type IAccountRepository interface {
	GetAllAccounts(int, int) ([]Account, error)
	GetAccountById(int64) (*Account, error)
	CreateAccount(dto.CreateAccountDto) (*Account, error)
	UpdateAccount(int64, dto.UpdateAccountDto) error
	DeleteAccount(int64) error
}

func (a Account) CreateAccountResponse() dto.AccountDto {
	return dto.AccountDto{
		Id:        a.Id,
		Owner:     a.Owner,
		Balance:   a.Balance,
		Currency:  a.Currency,
		CreatedAt: a.CreatedAt,
	}
}
