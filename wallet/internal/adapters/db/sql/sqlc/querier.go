// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddCardBalance(ctx context.Context, db DBTX, arg AddCardBalanceParams) error
	CreateAccount(ctx context.Context, db DBTX, arg CreateAccountParams) error
	CreateCard(ctx context.Context, db DBTX, arg CreateCardParams) error
	CreateTransaction(ctx context.Context, db DBTX, arg CreateTransactionParams) error
	CreateUser(ctx context.Context, db DBTX, externalID uuid.UUID) error
	DeleteAccount(ctx context.Context, db DBTX, id int32) error
	DeleteAccountCards(ctx context.Context, db DBTX, accountID int32) error
	DeleteCard(ctx context.Context, db DBTX, number string) error
	GetAccount(ctx context.Context, db DBTX, id int32) (Account, error)
	GetAccountCards(ctx context.Context, db DBTX, accountID int32) ([]Card, error)
	GetAccounts(ctx context.Context, db DBTX, userID int32) ([]Account, error)
	GetCard(ctx context.Context, db DBTX, number string) (Card, error)
	GetCardBalance(ctx context.Context, db DBTX, number string) (string, error)
	GetTransaction(ctx context.Context, db DBTX, id uuid.UUID) (GetTransactionRow, error)
	GetTransactions(ctx context.Context, db DBTX, arg GetTransactionsParams) ([]GetTransactionsRow, error)
	GetUserByExternalID(ctx context.Context, db DBTX, externalID uuid.UUID) (int32, error)
	GetUserCards(ctx context.Context, db DBTX, userID int32) ([]GetUserCardsRow, error)
	SubtractCardBalance(ctx context.Context, db DBTX, arg SubtractCardBalanceParams) error
}

var _ Querier = (*Queries)(nil)
