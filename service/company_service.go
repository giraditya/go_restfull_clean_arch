package service

import (
	"context"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
)

type CompanyService interface {
	Save(ctx context.Context, request request.CompanyCreateRequest) response.CompanyResponse
	Update(ctx context.Context, request request.CompanyUpdateRequest) response.CompanyResponse
	FindByID(ctx context.Context, companyID int) response.CompanyResponse
	FindAll(ctx context.Context) []response.CompanyResponse
}
