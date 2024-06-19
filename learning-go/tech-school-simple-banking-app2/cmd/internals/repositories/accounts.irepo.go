package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/errs"
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
	GetAllAccounts(int, int) ([]Account, *errs.AppError)
	CreateAccount(dto.CreateAccountDto) (*Account, *errs.AppError)
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
