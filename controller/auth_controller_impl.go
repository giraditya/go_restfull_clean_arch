package controller

import (
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) RequestToken(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := request.UserGenerateAuthKeyRequest{}
	helper.ReadFromRequestBody(r, &userRequest)

	tokenResponse := controller.AuthService.GenerateAuthKey(r.Context(), userRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tokenResponse,
	}
	helper.WriteToResponseBody(rw, webResponse)
}
