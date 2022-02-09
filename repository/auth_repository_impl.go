package repository

import (
	"context"
	"database/sql"
	"fmt"
	"giricorp/belajar-go-restfull-api/helper"
	"math/rand"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetAuthKey(ctx context.Context, tx *sql.Tx, authKey string) interface{} {
	SQL := "SELECT authKey FROM user WHERE authKey = ?"
	rows, err := tx.QueryContext(ctx, SQL, authKey)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return authKey
	} else {
		return nil
	}
}

func (repository *AuthRepositoryImpl) GenerateAuthKey(ctx context.Context, tx *sql.Tx, username string, length int) string {
	var letterBytes = fmt.Sprintf("69HASDHDH&GIRICORP%s", username)
	authKey := make([]byte, length)
	for i := range authKey {
		authKey[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	SQL := "UPDATE user SET authKey = ? WHERE username = ?"
	_, err := tx.ExecContext(ctx, SQL, string(authKey), username)
	helper.PanicIfError(err)

	return string(authKey)
}

func (repository *AuthRepositoryImpl) CredentialsValid(ctx context.Context, tx *sql.Tx, username string, password string) bool {
	SQL := "SELECT user.id, user.name, user.address, user.username, company.name FROM user LEFT JOIN company ON company.ID = user.companyID WHERE user.username = ? AND user.password = ?"
	rows, err := tx.QueryContext(ctx, SQL, username, password)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}
