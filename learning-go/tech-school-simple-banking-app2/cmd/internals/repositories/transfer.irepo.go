package repositories

import (
	"banking_app2/cmd/internals/dto"
	"time"
)

type Transfer struct {
	Id            int64     `db:"id"`
	FromAccountId int64     `db:"from_account_id"`
	ToAccountId   int64     `db:"to_account_id"`
	Amount        int64     `db:"amount"`
	CreatedAt     time.Time `db:"created_at"`
}

type ITransferRepo interface {
	TransferAmount(dto.TransferAmountDto) (Transfer, error)
}
