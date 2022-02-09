package repository

import (
	"context"
	"database/sql"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
