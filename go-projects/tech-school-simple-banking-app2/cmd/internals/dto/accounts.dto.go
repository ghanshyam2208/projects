package dto

import (
	"time"

	"github.com/labstack/echo"
)

type AccountDto struct {
	Id        int64     `json:"customer_id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type SuccessFullApiResponseDto struct {
	ctx  echo.Context
	code int
	data interface{}
}

// Define a struct to receive the request data
type CreateAccountDto struct {
	Owner    string `json:"owner" validate:"required"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency" validate:"required,oneof=USD INR EUR AUS JPY GBP"`
}

type UpdateAccountDto struct {
	Id       int64   `json:"id" validate:"required"`
	Owner    *string // Pointers to handle optional field
	Currency *string
	Balance  *int64
}