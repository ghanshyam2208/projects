package dto

import "time"

type CreateUserDto struct {
	Username string `json:"username"  validate:"required,min=3,max=20"`
	Password string `json:"password"  validate:"required,min=8,max=20,customPasswordValidationHandler"`
	FullName string `json:"full_name"  validate:"required"`
	Email    string `json:"email"  validate:"required"`
}

type UserResp struct {
	Id        int64     `json:"user_id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
