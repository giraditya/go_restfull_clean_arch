package service

import (
	"context"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
)

type AuthService interface {
	IsUserAuthenticated(ctx context.Context, authKey string) bool
	GenerateAuthKey(ctx context.Context, request request.UserGenerateAuthKeyRequest) response.UserAuthKeyResponse
	CredentialsValid(ctx context.Context, request request.UserGenerateAuthKeyRequest) bool
}
