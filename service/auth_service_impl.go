package service

import (
	"context"
	"database/sql"
	"giricorp/belajar-go-restfull-api/exception"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
	"giricorp/belajar-go-restfull-api/repository"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) IsUserAuthenticated(ctx context.Context, authKey string) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result := service.AuthRepository.GetAuthKey(ctx, tx, authKey)

	switch result.(type) {
	default:
		return false
	case string:
		return true
	}
}

func (service *AuthServiceImpl) GenerateAuthKey(ctx context.Context, request request.UserGenerateAuthKeyRequest) response.UserAuthKeyResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	credentialsValid, err := service.AuthRepository.CredentialsValid(ctx, tx, request.Username, request.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	var authKey string
	if credentialsValid {
		authKey = service.AuthRepository.GenerateAuthKey(ctx, tx, request.Username, 12)
	}

	return helper.ToAuthKeyResponse(authKey)
}

func (service *AuthServiceImpl) CredentialsValid(ctx context.Context, request request.UserGenerateAuthKeyRequest) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.AuthRepository.CredentialsValid(ctx, tx, request.Username, request.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	return result
}
