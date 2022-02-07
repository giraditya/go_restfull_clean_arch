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

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validator *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validator,
	}
}

func (service *ProductServiceImpl) Save(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name: request.Name,
		Size: request.Size,
		Category: domain.Category{
			ID: request.CategoryID,
		},
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Size = request.Size
	product.Category = domain.Category{
		ID: request.CategoryID,
	}
	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productID int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, productID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product.ID)
}

func (service *ProductServiceImpl) FindByID(ctx context.Context, productID int) response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindByID(ctx, tx, productID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}
