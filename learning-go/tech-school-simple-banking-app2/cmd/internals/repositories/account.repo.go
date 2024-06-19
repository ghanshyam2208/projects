package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/errs"
	"banking_app2/cmd/utils/logger"
	"time"

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
		return nil, errs.NewInternalServerError("Error while getting all accounts: " + err.Error())
	}
	return accounts, nil
}

func (d *AccountRepositoryDB) connectDB() {
	// connect to database
	sqlDb, err := sqlx.Connect("postgres", "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable")
	if err != nil {
		logger.Error("Error while connecting to database: " + err.Error())
		// panic(err) // this will stop the server, as of now lets not stop server if not able to connect to the server
	}
	d.sqlxClient = sqlDb
	logger.Info("Successfully connected to the database")
}

func NewAccountsRepo() *AccountRepositoryDB {
	repo := &AccountRepositoryDB{}
	repo.connectDB()
	return repo
}

func (d *AccountRepositoryDB) CreateAccount(createAccountDto dto.CreateAccountDto) (*Account, *errs.AppError) {
	account := Account{
		Owner:     createAccountDto.Owner,
		Balance:   createAccountDto.Balance,
		Currency:  createAccountDto.Currency,
		CreatedAt: time.Now(),
	}

	// Insert the account into the database
	query := `INSERT INTO accounts (owner, balance, currency, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := d.sqlxClient.QueryRowx(query, account.Owner, account.Balance, account.Currency, account.CreatedAt).Scan(&account.Id)

	if err != nil {
		logger.Error("Error while creating account: " + err.Error())
		return nil, errs.NewInternalServerError("Error while creating account: " + err.Error())
	}

	return &account, nil
}
