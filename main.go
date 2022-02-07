package main

import (
	"giricorp/belajar-go-restfull-api/app"
	"giricorp/belajar-go-restfull-api/controller"
	"giricorp/belajar-go-restfull-api/exception"
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/middleware"
	"giricorp/belajar-go-restfull-api/repository"
	"giricorp/belajar-go-restfull-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validator := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validator)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryID", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Save)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productID", productController.FindByID)
	router.POST("/api/products", productController.Save)
	router.PUT("/api/products/:productID", productController.Update)
	router.DELETE("/api/products/:productID", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
