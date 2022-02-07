package helper

import (
	"giricorp/belajar-go-restfull-api/model/api/response"
	"giricorp/belajar-go-restfull-api/model/domain"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []response.CategoryResponse {
	var categoryReponses []response.CategoryResponse
	for _, category := range categories {
		categoryReponses = append(categoryReponses, ToCategoryResponse(category))
	}
	return categoryReponses
}

func ToProductResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Size:     product.Size,
		Category: product.Category.Name,
	}
}

func ToProductResponses(products []domain.Product) []response.ProductResponse {
	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
