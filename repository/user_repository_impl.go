package repository

import (
	"context"
	"database/sql"
	"errors"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(name, address, username, password) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Address, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.ID = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET name = ?, address = ?, username = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Address, user.Username, user.Password, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID int) (domain.User, error) {
	SQL := "SELECT user.id, user.name, user.address, user.username, company.name FROM user LEFT JOIN company ON company.id = user.companyID WHERE user.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)
	defer rows.Close()
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Username, &user.Company.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT user.id, user.name, user.address, user.username, company.name FROM user LEFT JOIN company ON user.companyID = company.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Username, &user.Company.Name)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
