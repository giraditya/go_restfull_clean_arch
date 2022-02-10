package middleware

import (
	"giricorp/belajar-go-restfull-api/exception"
	"giricorp/belajar-go-restfull-api/service"
	"net/http"
)

type AuthMiddleware struct {
	Handler     http.Handler
	AuthService service.AuthService
	Exception   []ExceptionHandler
}

type ExceptionHandler struct {
	HandlerName string
	Method      string
}

func NewAuthMiddleware(handler http.Handler, authService service.AuthService, exceptionHandler []ExceptionHandler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler:     handler,
		AuthService: authService,
		Exception:   exceptionHandler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	authKey := r.Header.Get("X-api-Key")
	isUserAuthenticated := middleware.AuthService.IsUserAuthenticated(r.Context(), authKey)
	isExceptionHandler := false

	if !isUserAuthenticated {
		for _, x := range middleware.Exception {
			if x.HandlerName == r.URL.Path {
				if x.Method == r.Method {
					isExceptionHandler = true
				}
			}
		}
	}

	if isUserAuthenticated || isExceptionHandler {
		middleware.Handler.ServeHTTP(rw, r)
	} else {
		panic(exception.NewUnauthorizedError("credentials is invalid"))
	}
}
