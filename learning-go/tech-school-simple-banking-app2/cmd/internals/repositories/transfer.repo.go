package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransferRepositoryDb struct {
	sqlxClient *sqlx.DB
	appConfigs *configs.Config
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

	return &Transfer{
		Id:            1,
		FromAccountId: payload.FromAccountId,
		ToAccountId:   payload.ToAccountId,
		Amount:        payload.Amount,
		CreatedAt:     time.Now(),
	}, nil
}

func NewTransferRepo(configs *configs.Config) *TransferRepositoryDb {
	repo := &TransferRepositoryDb{
		appConfigs: configs,
	}
	repo.connectDB2()
	return repo
}

func (t *TransferRepositoryDb) connectDB2() {
	// connect to database
	sqlDb, err := sqlx.Connect("postgres", t.appConfigs.PostgresConnStr)
	if err != nil {
		logger.Error("Error while connecting to database: " + err.Error())
		// panic(err) // this will stop the server, as of now lets not stop server if not able to connect to the server
	}
	t.sqlxClient = sqlDb
	logger.Info("Successfully connected to the database")
}
