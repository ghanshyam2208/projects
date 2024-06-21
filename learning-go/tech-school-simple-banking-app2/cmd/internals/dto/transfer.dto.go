package dto

type TransferAmountDto struct {
	FromAccountId int64 `json:"from_account_id"  validate:"required"`
	ToAccountId   int64 `json:"to_account_id"  validate:"required"`
	Amount        int64 `json:"amount"  validate:"required"`
}
