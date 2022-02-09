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
	db := app.NewDBMain()

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validator)
	authController := controller.NewAuthController(authService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validator)
	productController := controller.NewProductController(productService)
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validator)
	userController := controller.NewUserController(userService)
	companyRepository := repository.NewCompanyRepository()
	companyService := service.NewCompanyService(companyRepository, db, validator)
	companyController := controller.NewCompanyController(companyService)

	router := httprouter.New()

	// Auth
	router.POST("/api/auth/requesttoken", authController.RequestToken)

	// User
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userID", userController.FindByID)
	router.POST("/api/users", userController.Save)
	router.PUT("/api/users/:userID", userController.Update)

	// Company
	router.GET("/api/company", companyController.FindAll)
	router.GET("/api/company/:companyID", companyController.FindByID)
	router.POST("/api/company", companyController.Save)
	router.PUT("/api/company/:companyID", companyController.Update)

	// Categories
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryID", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Save)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)

	// Products
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productID", productController.FindByID)
	router.POST("/api/products", productController.Save)
	router.PUT("/api/products/:productID", productController.Update)
	router.DELETE("/api/products/:productID", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	exceptionHandler := []middleware.ExceptionHandler{
		{
			HandlerName: "/api/users",
			Method:      "POST",
		},
		{
			HandlerName: "/api/auth/requesttoken",
			Method:      "POST",
		},
	}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router, authService, exceptionHandler),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
