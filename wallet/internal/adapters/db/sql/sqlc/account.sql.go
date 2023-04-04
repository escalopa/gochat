// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (user_id, currency_id, name)
VALUES ($1, $2, $3)
`

type CreateAccountParams struct {
	UserID     int32  `db:"user_id" json:"user_id"`
	CurrencyID int16  `db:"currency_id" json:"currency_id"`
	Name       string `db:"name" json:"name"`
}

func (q *Queries) CreateAccount(ctx context.Context, db DBTX, arg CreateAccountParams) error {
	_, err := db.ExecContext(ctx, createAccount, arg.UserID, arg.CurrencyID, arg.Name)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE
FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, db DBTX, id int32) error {
	_, err := db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_id, name, balance, currency_id
FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, db DBTX, id int32) (Account, error) {
	row := db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.CurrencyID,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT id, user_id, name, balance, currency_id
FROM accounts
WHERE user_id = $1
`

func (q *Queries) GetAccounts(ctx context.Context, db DBTX, userID int32) ([]Account, error) {
	rows, err := db.QueryContext(ctx, getAccounts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Balance,
			&i.CurrencyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
