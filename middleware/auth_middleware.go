package middleware

import (
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api"
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

	for _, x := range middleware.Exception {
		if x.HandlerName == r.URL.Path {
			if x.Method == r.Method {
				isUserAuthenticated = true
			}
		}
	}

	if isUserAuthenticated {
		middleware.Handler.ServeHTTP(rw, r)
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusUnauthorized)

		webResponse := api.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(rw, webResponse)
	}
}
