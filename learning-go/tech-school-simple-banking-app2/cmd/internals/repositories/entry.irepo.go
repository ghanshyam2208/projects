package repositories

import "time"

type Entry struct {
	Id        int64     `db:"id"`
	AccountId int64     `db:"account_id"`
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
