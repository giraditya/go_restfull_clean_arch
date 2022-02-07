package repository

import (
	"context"
	"database/sql"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, productID int)
	FindByID(ctx context.Context, tx *sql.Tx, productID int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
