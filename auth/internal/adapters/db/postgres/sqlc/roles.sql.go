// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: roles.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createRole = `-- name: CreateRole :exec
INSERT INTO roles (name)
VALUES ($1)
`

func (q *Queries) CreateRole(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createRole, name)
	return err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT id, name
FROM roles
WHERE name = $1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserRoles = `-- name: GetUserRoles :many
SELECT r.name
FROM user_roles ur
       JOIN roles r on r.id = ur.role_id
WHERE ur.user_id = $1
`

func (q *Queries) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getUserRoles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const grantRoleToUser = `-- name: GrantRoleToUser :exec
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2)
`

type GrantRoleToUserParams struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	RoleID int32     `db:"role_id" json:"role_id"`
}

func (q *Queries) GrantRoleToUser(ctx context.Context, arg GrantRoleToUserParams) error {
	_, err := q.db.ExecContext(ctx, grantRoleToUser, arg.UserID, arg.RoleID)
	return err
}

const revokeRoleFromUser = `-- name: RevokeRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = $1
  AND role_id = $2
`

type RevokeRoleFromUserParams struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	RoleID int32     `db:"role_id" json:"role_id"`
}

func (q *Queries) RevokeRoleFromUser(ctx context.Context, arg RevokeRoleFromUserParams) error {
	_, err := q.db.ExecContext(ctx, revokeRoleFromUser, arg.UserID, arg.RoleID)
	return err
}