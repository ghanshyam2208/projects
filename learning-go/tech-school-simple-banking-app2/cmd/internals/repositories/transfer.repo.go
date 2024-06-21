package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"
	"time"
)

type TransferRepositoryDb struct {
	RepositoryDB
}

func (t *TransferRepositoryDb) TransferAmount(payload dto.TransferAmountDto) (*Transfer, error) {
	tx, stdErr := t.sqlxClient.Beginx()
	if stdErr != nil {
		return nil, stdErr
	}

	defer func() {
		if stdErr != nil {
			_ = tx.Rollback() // If there's an error during a rollback, it typically indicates an issue that might not be recoverable or actionable within the current context of handling another error
		} else {
			stdErr = tx.Commit()
		}
	}()

	// Decrease the balance of the fromAccount
	query := "UPDATE accounts SET balance = balance - $1 WHERE id = $2"
	_, stdErr = tx.Exec(query, payload.Amount, payload.FromAccountId)
	if stdErr != nil {
		return nil, stdErr
	}

	// increase the balance to toAccount
	query = "UPDATE accounts SET balance = balance + $1 WHERE id = $2"
	_, stdErr = tx.Exec(query, payload.Amount, payload.ToAccountId)
	if stdErr != nil {
		return nil, stdErr
	}

	// Insert the transaction entry for the fromAccount
	entry := Entry{
		AccountId: payload.FromAccountId,
		Amount:    -payload.Amount,
		CreatedAt: time.Now(),
	}
	query = "INSERT INTO entries (account_id, amount, created_at) VALUES ($1, $2, $3) RETURNING id"
	stdErr = tx.QueryRowx(query, entry.AccountId, entry.Amount, entry.CreatedAt).Scan(&entry.Id)
	if stdErr != nil {
		return nil, stdErr
	}

	// Insert the transaction entry for the fromAccount
	entry = Entry{
		AccountId: payload.ToAccountId,
		Amount:    +payload.Amount,
		CreatedAt: time.Now(),
	}
	query = "INSERT INTO entries (account_id, amount, created_at) VALUES ($1, $2, $3) RETURNING id"
	stdErr = tx.QueryRowx(query, entry.AccountId, entry.Amount, entry.CreatedAt).Scan(&entry.Id)
	if stdErr != nil {
		return nil, stdErr
	}

	transfer := Transfer{
		FromAccountId: payload.FromAccountId,
		ToAccountId:   payload.ToAccountId,
		Amount:        payload.Amount,
		CreatedAt:     time.Now(),
	}

	// Insert the transfer into the database
	query = "INSERT INTO transfers (from_account_id, to_account_id, amount, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	stdErr = tx.QueryRowx(query, transfer.FromAccountId, transfer.ToAccountId, transfer.Amount, transfer.CreatedAt).Scan(&transfer.Id)
	if stdErr != nil {
		return nil, stdErr
	}

	return &transfer, nil
}

func NewTransferRepo(configs *configs.Config) *TransferRepositoryDb {
	repo := &TransferRepositoryDb{}
	repo.appConfigs = configs
	if err := repo.connectDB(); err != nil {
		logger.Error("Error connecting to database for accounts repository: " + err.Error())
		// Handle error as needed, e.g., return nil or panic
	}
	return repo
}
