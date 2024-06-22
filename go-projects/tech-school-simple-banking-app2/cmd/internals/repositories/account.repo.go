package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type AccountRepositoryDB struct {
	RepositoryDB
}

func (d *AccountRepositoryDB) GetAllAccounts(page int, pageSize int) ([]Account, error) {
	accounts := []Account{}
	offset := (page - 1) * pageSize
	query := "SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2"
	err := d.sqlxClient.Select(&accounts, query, pageSize, offset)
	if err != nil {
		logger.Error("Error while getting all accounts: " + err.Error())
		return nil, err
	}
	return accounts, nil
}

func (d *AccountRepositoryDB) GetAccountById(customer_id int64) (*Account, error) {
	accounts := Account{}
	query := "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1 "
	err := d.sqlxClient.Get(&accounts, query, customer_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No records found: " + err.Error())
			return nil, errors.New("no records found ")
		}
		logger.Error("Error while getting all accounts: " + err.Error())
		return nil, err
	}
	return &accounts, nil
}

// func (d *AccountRepositoryDB) connectDB() {
// 	// connect to database
// 	sqlDb, err := sqlx.Connect("postgres", d.appConfigs.PostgresConnStr)
// 	if err != nil {
// 		logger.Error("Error while connecting to database: " + err.Error())
// 		// panic(err) // this will stop the server, as of now lets not stop server if not able to connect to the server
// 	}
// 	d.sqlxClient = sqlDb
// 	logger.Info("Successfully connected to the database")
// }

func NewAccountsRepo(configs *configs.Config) *AccountRepositoryDB {
	repo := &AccountRepositoryDB{}
	repo.appConfigs = configs
	if err := repo.connectDB(); err != nil {
		logger.Error("Error connecting to database for accounts repository: " + err.Error())
		// Handle error as needed, e.g., return nil or panic
	}
	return repo
}

func (d *AccountRepositoryDB) CreateAccount(createAccountDto dto.CreateAccountDto) (*Account, error) {
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
		return nil, err
	}

	return &account, nil
}

func (d *AccountRepositoryDB) UpdateAccount(id int64, updateAccountDto dto.UpdateAccountDto) error {
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
		logger.Error("No fields to update")
		return errors.New("no fields to update")
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
		return stdErr
	}

	return nil
}

func (d *AccountRepositoryDB) DeleteAccount(id int64) error {
	query := "DELETE FROM accounts WHERE id = $1"
	result, stdErr := d.sqlxClient.Exec(query, id)
	if stdErr != nil {
		logger.Error("Error while deleting account: " + stdErr.Error())
		return stdErr
	}

	rowsAffected, stdErr := result.RowsAffected()
	if stdErr != nil {
		logger.Error("Error while checking rows affected: " + stdErr.Error())
		return stdErr
	}

	if rowsAffected == 0 {
		logger.Error("No records found for deletion")
		return errors.New("no records found")
	}

	return nil
}

func (d *AccountRepositoryDB) CleanAccounts() error {
	query := "DELETE FROM accounts"
	_, stdErr := d.sqlxClient.Exec(query)
	if stdErr != nil {
		logger.Error("Error while deleting account: " + stdErr.Error())
		return stdErr
	}

	return nil
}
