package repository

import (
	"context"
	"database/sql"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type CompanyRepository interface {
	Save(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company
	Update(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company
	FindByID(ctx context.Context, tx *sql.Tx, companyID int) (domain.Company, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Company
}
