// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateRole(ctx context.Context, name string) error
	CreateSession(ctx context.Context, arg CreateSessionParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteSessionByID(ctx context.Context, id uuid.UUID) (int64, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) (int64, error)
	GetRoleByName(ctx context.Context, name string) (Role, error)
	GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserDevices(ctx context.Context, userID uuid.UUID) ([]GetUserDevicesRow, error)
	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]string, error)
	GetUserSessions(ctx context.Context, userID uuid.UUID) ([]Session, error)
	GrantRoleToUser(ctx context.Context, arg GrantRoleToUserParams) (int64, error)
	RevokeRoleFromUser(ctx context.Context, arg RevokeRoleFromUserParams) (int64, error)
	SetSessionIsBlocked(ctx context.Context, arg SetSessionIsBlockedParams) (int64, error)
	UpdateSessionRefreshToken(ctx context.Context, arg UpdateSessionRefreshTokenParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
