package repository

import (
	"context"
	"database/sql"
)

type AuthRepository interface {
	GetAuthKey(ctx context.Context, tx *sql.Tx, authKey string) interface{}
	GenerateAuthKey(ctx context.Context, tx *sql.Tx, username string, length int) string
	CredentialsValid(ctx context.Context, tx *sql.Tx, username string, password string) bool
}
