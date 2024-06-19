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
