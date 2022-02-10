package exception

import (
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(rw http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(rw, r, err) {
		return
	}
	if validationErrors(rw, r, err) {
		return
	}
	if unauthorizedError(rw, r, err) {
		return
	}
	internalServerError(rw, r, err)
}

func validationErrors(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)

		webResponse := api.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)

		webResponse := api.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not found",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedError(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusUnauthorized)

		webResponse := api.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(rw http.ResponseWriter, r *http.Request, err interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)

	webResponse := api.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal server error",
		Data:   err,
	}
	helper.WriteToResponseBody(rw, webResponse)
}
