package repositories

import (
	"banking_app2/cmd/utils/errs"
	"banking_app2/cmd/utils/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AccountRepositoryDB struct {
	sqlxClient *sqlx.DB
}

func (d *AccountRepositoryDB) GetAllAccounts() ([]Account, *errs.AppError) {
	accounts := []Account{}
	query := "SELECT id, owner, balance, currency, created_at FROM accounts"
	err := d.sqlxClient.Select(&accounts, query)
	if err != nil {
		logger.Error("Error while getting all accounts: " + err.Error())
		return nil, errs.NewInternalServerError()
	}
	return accounts, nil
}

func (d *AccountRepositoryDB) connectDB() {
	// connect to database
	sqlDb, err := sqlx.Connect("postgres", "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable")
	if err != nil {
		logger.Error("Error while connecting to database: " + err.Error())
		panic(err)
	}
	d.sqlxClient = sqlDb
	logger.Info("Successfully connected to the database")
}

func NewAccountsRepo() *AccountRepositoryDB {
	repo := &AccountRepositoryDB{}
	repo.connectDB()
	return repo
}
