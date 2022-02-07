package service

import (
	"context"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
)

type ProductService interface {
	Save(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	Delete(ctx context.Context, productID int)
	FindByID(ctx context.Context, productID int) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
