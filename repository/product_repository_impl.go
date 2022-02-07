package repository

import (
	"context"
	"database/sql"
	"errors"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO product(name, size, categoryID) VALUES (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Size, product.Category.ID)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.ID = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE product SET (name, size, categoryID) VALUES (?,?,?) WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Size, product.Category.ID, product.ID)
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productID int) {
	SQL := "DELETE FROM product WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productID)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, productID int) (domain.Product, error) {
	SQL := "SELECT product.id, product.name, product.size, category.name FROM product LEFT JOIN category ON product.categoryID = category.id WHERE product.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productID)
	helper.PanicIfError(err)
	defer rows.Close()
	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Size, &product.Category.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT product.id, product.name, product.size, category.name FROM product LEFT JOIN category ON product.categoryID = category.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Size, &product.Category.Name)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}
