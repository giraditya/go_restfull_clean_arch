package service

import (
	"context"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
)

type UserService interface {
	Save(ctx context.Context, request request.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse
	FindByID(ctx context.Context, userID int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}
