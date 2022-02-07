package service

import (
	"context"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
)

type CategoryService interface {
	Save(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryID int)
	FindByID(ctx context.Context, categoryID int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
