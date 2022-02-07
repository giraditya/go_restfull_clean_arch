package repository

import (
	"context"
	"database/sql"
	"errors"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (respository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.ID = int(id)
	return category
}

func (respository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	helper.PanicIfError(err)

	return category
}

func (respository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryID int) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryID)
	helper.PanicIfError(err)
}

func (respository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (respository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
