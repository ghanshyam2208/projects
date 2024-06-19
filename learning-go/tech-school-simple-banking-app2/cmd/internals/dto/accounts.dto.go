package dto

import "time"

type AccountDto struct {
	Id        int64     `json:"customer_id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
