package service

import (
	"context"
	"database/sql"
	"giricorp/belajar-go-restfull-api/exception"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/model/api/response"
	"giricorp/belajar-go-restfull-api/model/domain"
	"giricorp/belajar-go-restfull-api/repository"

	"github.com/go-playground/validator/v10"
)

type CompanyServiceImpl struct {
	CompanyRepository repository.CompanyRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCompanyService(companyRepository repository.CompanyRepository, DB *sql.DB, validate *validator.Validate) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: companyRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CompanyServiceImpl) Save(ctx context.Context, request request.CompanyCreateRequest) response.CompanyResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company := domain.Company{
		Code: request.Code,
		Name: request.Name,
	}

	company = service.CompanyRepository.Save(ctx, tx, company)

	return helper.ToCompanyResponse(company)
}

func (service *CompanyServiceImpl) Update(ctx context.Context, request request.CompanyUpdateRequest) response.CompanyResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company, err := service.CompanyRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	company.Name = request.Name
	company.Code = request.Code
	company = service.CompanyRepository.Update(ctx, tx, company)

	return helper.ToCompanyResponse(company)
}

func (service *CompanyServiceImpl) FindByID(ctx context.Context, companyID int) response.CompanyResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company, err := service.CompanyRepository.FindByID(ctx, tx, companyID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCompanyResponse(company)
}

func (service *CompanyServiceImpl) FindAll(ctx context.Context) []response.CompanyResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	companies := service.CompanyRepository.FindAll(ctx, tx)

	return helper.ToCompanyResponses(companies)
}
