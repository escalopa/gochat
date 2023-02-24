// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const changeNames = `-- name: ChangeNames :exec
UPDATE users
SET name     = coalesce($1, name),
    username = coalesce($2, name)
WHERE id = $3
`

type ChangeNamesParams struct {
	Name     sql.NullString `db:"name" json:"name"`
	Username sql.NullString `db:"username" json:"username"`
	ID       uuid.UUID      `db:"id" json:"id"`
}

func (q *Queries) ChangeNames(ctx context.Context, arg ChangeNamesParams) error {
	_, err := q.db.ExecContext(ctx, changeNames, arg.Name, arg.Username, arg.ID)
	return err
}

const changePassword = `-- name: ChangePassword :exec
UPDATE users
SET hashed_password = $2
WHERE id = $1
`

type ChangePasswordParams struct {
	ID             uuid.UUID `db:"id" json:"id"`
	HashedPassword string    `db:"hashed_password" json:"hashed_password"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) error {
	_, err := q.db.ExecContext(ctx, changePassword, arg.ID, arg.HashedPassword)
	return err
}

const changeUserEmail = `-- name: ChangeUserEmail :exec
UPDATE users
SET email       = $2,
    is_verified = false
WHERE id = $1
`

type ChangeUserEmailParams struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Email string    `db:"email" json:"email"`
}

func (q *Queries) ChangeUserEmail(ctx context.Context, arg ChangeUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, changeUserEmail, arg.ID, arg.Email)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, username, email, hashed_password)
VALUES ($1, $2, $3, $4, $5)
`

type CreateUserParams struct {
	ID             uuid.UUID `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Username       string    `db:"username" json:"username"`
	Email          string    `db:"email" json:"email"`
	HashedPassword string    `db:"hashed_password" json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
	)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, username, email, hashed_password, is_verified, created_at
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, username, email, hashed_password, is_verified, created_at
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, name, username, email, hashed_password, is_verified, created_at
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.IsVerified,
		&i.CreatedAt,
	)
	return i, err
}

const setUserIsVerified = `-- name: SetUserIsVerified :exec
UPDATE users
SET is_verified = $1
WHERE id = $1
`

func (q *Queries) SetUserIsVerified(ctx context.Context, isVerified bool) error {
	_, err := q.db.ExecContext(ctx, setUserIsVerified, isVerified)
	return err
}
