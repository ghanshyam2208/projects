package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/errs"
	"banking_app2/cmd/utils/logger"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AccountRepositoryDB struct {
	sqlxClient *sqlx.DB
}

func (d *AccountRepositoryDB) GetAllAccounts(page int, pageSize int) ([]Account, *errs.AppError) {
	accounts := []Account{}
	offset := (page - 1) * pageSize
	query := "SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2"
	err := d.sqlxClient.Select(&accounts, query, pageSize, offset)
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

func (d *AccountRepositoryDB) UpdateAccount(id int64, updateAccountDto dto.UpdateAccountDto) *errs.AppError {
	var query strings.Builder
	var args []interface{}

	query.WriteString("UPDATE accounts SET ")

	// Add each field conditionally to the set clause and args slice
	var setClauses []string

	if updateAccountDto.Owner != nil {
		setClauses = append(setClauses, "owner = $1")
		args = append(args, *updateAccountDto.Owner)
	}
	if updateAccountDto.Currency != nil {
		setClauses = append(setClauses, "currency = $2")
		args = append(args, *updateAccountDto.Currency)
	}
	if updateAccountDto.Balance != nil {
		setClauses = append(setClauses, "balance = $3")
		args = append(args, *updateAccountDto.Balance)
	}

	// Join the set clauses with commas
	if len(setClauses) == 0 {
		return errs.NewInternalServerError("No fields to update")
	}

	query.WriteString(strings.Join(setClauses, ", "))
	query.WriteString(" WHERE id = $4")
	args = append(args, id)

	queryString := query.String()
	logger.Debug("query: " + queryString)
	fmt.Println(queryString)
	fmt.Println(args...)

	// Execute the update query
	_, stdErr := d.sqlxClient.Exec(queryString, args...)
	if stdErr != nil {
		logger.Error("Error while updating account: " + stdErr.Error())
		return errs.NewInternalServerError("Error while updating account: " + stdErr.Error())
	}

	return nil
}
