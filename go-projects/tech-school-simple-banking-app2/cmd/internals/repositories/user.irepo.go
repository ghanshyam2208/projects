package repositories

import (
	"banking_app2/cmd/internals/dto"
	"time"
)

type User struct {
	Id                int64     `db:"id"`
	Username          string    `db:"username"`
	HashedPassword    string    `db:"hashed_password"`
	FullName          string    `db:"full_name"`
	Email             string    `db:"email"`
	PasswordChangedAt time.Time `db:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at"`
}

type IUserRepo interface {
	GetAllUsers(int, int) ([]User, error)
	GetUserById(int64) (*User, error)
	CreateUser(dto.CreateUserDto) (*User, error)
	CheckPassword(string, string) error
}

func (u User) CreateUserResponse() dto.UserResp {
	return dto.UserResp{
		Id:        u.Id,
		Username:  u.Username,
		FullName:  u.FullName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
