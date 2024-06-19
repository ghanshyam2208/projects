package repositories

import (
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
	GetAllAccounts() ([]Account, *errs.AppError)
}
