// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
)

type Accounts struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entries struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can positive or negative
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Tranfers struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// positive
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}
