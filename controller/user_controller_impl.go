package controller

import (
	"giricorp/belajar-go-restfull-api/helper"
	"giricorp/belajar-go-restfull-api/model/api"
	"giricorp/belajar-go-restfull-api/model/api/request"
	"giricorp/belajar-go-restfull-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := controller.UserService.Save(r.Context(), userCreateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *UserControllerImpl) Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userUpdateRequest := request.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userID := params.ByName("userID")
	id, err := strconv.Atoi(userID)
	helper.PanicIfError(err)

	userUpdateRequest.ID = id

	userResponse := controller.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *UserControllerImpl) FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("userID")
	ID, err := strconv.Atoi(userID)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindByID(r.Context(), ID)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *UserControllerImpl) FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(r.Context())
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
