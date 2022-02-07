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

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Save(r.Context(), categoryCreateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryID := params.ByName("categoryID")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryUpdateRequest.ID = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) Delete(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryID")
	ID, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), ID)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryID")
	ID, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(r.Context(), ID)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
