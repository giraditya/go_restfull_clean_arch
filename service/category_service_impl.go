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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.ID)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID int) response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
