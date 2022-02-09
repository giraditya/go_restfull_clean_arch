package repository

import (
	"context"
	"database/sql"
	"errors"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type CompanyRepositoryImpl struct {
}

func NewCompanyRepository() CompanyRepository {
	return &CompanyRepositoryImpl{}
}

func (repository *CompanyRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company {
	SQL := "INSERT INTO company(name, code) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, SQL, company.Name, company.Code)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	company.ID = int(id)
	return company
}

func (repository *CompanyRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company {
	SQL := "UPDATE company SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, company.Name, company.ID)
	helper.PanicIfError(err)

	return company
}

func (repository *CompanyRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, companyID int) (domain.Company, error) {
	SQL := "SELECT id, name, code FROM company WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, companyID)
	helper.PanicIfError(err)
	defer rows.Close()

	company := domain.Company{}
	if rows.Next() {
		err := rows.Scan(&company.ID, &company.Code, &company.Name)
		helper.PanicIfError(err)
		return company, nil
	} else {
		return company, errors.New("company not found")
	}
}

func (repository *CompanyRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Company {
	SQL := "SELECT id, name, code FROM company"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var companies []domain.Company
	for rows.Next() {
		company := domain.Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Code)
		helper.PanicIfError(err)
		companies = append(companies, company)
	}
	return companies
}
