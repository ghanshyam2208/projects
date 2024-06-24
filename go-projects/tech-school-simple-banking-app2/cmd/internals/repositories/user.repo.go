package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"
	"database/sql"
	"errors"
	"time"
)

type DefaultUserRepoDb struct {
	RepositoryDB
}

func (d *DefaultUserRepoDb) GetAllUsers(page int, pageSize int) ([]User, error) {
	return nil, nil
}

func (d *DefaultUserRepoDb) GetUserById(userId int64) (*User, error) {
	return nil, nil
}

func (d *DefaultUserRepoDb) CreateUser(userPayload dto.CreateUserDto) (*User, error) {
	hashedPassword, stdErr := helpers.HashPassword(userPayload.Password)
	if stdErr != nil {
		return nil, stdErr
	}

	user := User{
		Id:                1,
		Username:          userPayload.Username,
		HashedPassword:    hashedPassword,
		FullName:          userPayload.FullName,
		Email:             userPayload.Email,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}

	query := `INSERT INTO users (username, hashed_password, full_name, email, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stdErr = d.sqlxClient.QueryRowx(query, user.Username, user.HashedPassword, user.FullName, user.Email, user.CreatedAt).Scan(&user.Id)
	if stdErr != nil {
		logger.Error("Error while creating account: " + stdErr.Error())
		return nil, stdErr
	}

	return &user, nil
}

func (d *DefaultUserRepoDb) CheckPassword(username string, password string) error {
	user := User{}
	query := "SELECT username, hashed_password FROM users WHERE username = $1"
	stdErr := d.sqlxClient.Get(&user, query, username)
	if stdErr != nil {
		if stdErr == sql.ErrNoRows {
			logger.Error("No records found: " + stdErr.Error())
			return errors.New("no records found ")
		}
		logger.Error("Error while getting all accounts: " + stdErr.Error())
		return stdErr
	}
	return helpers.ComparePassword(user.HashedPassword, password)
}

func NewUserRepo(configs *configs.Config) *DefaultUserRepoDb {
	repo := &DefaultUserRepoDb{}
	repo.appConfigs = configs
	if err := repo.connectDB(); err != nil {
		logger.Error("Error connecting to database for users repository: " + err.Error())
	}
	return repo
}
