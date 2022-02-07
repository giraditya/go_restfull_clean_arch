package middleware

import (
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-api-Key")

	if apiKey == "SECRET" {
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
