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

type CompanyControllerImpl struct {
	CompanyService service.CompanyService
}

func NewCompanyController(companyService service.CompanyService) CompanyController {
	return &CompanyControllerImpl{
		CompanyService: companyService,
	}
}

func (controller *CompanyControllerImpl) Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	companyCreateRequest := request.CompanyCreateRequest{}
	helper.ReadFromRequestBody(r, &companyCreateRequest)

	companyResponse := controller.CompanyService.Save(r.Context(), companyCreateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   companyResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CompanyControllerImpl) Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	companyUpdateRequest := request.CompanyUpdateRequest{}
	helper.ReadFromRequestBody(r, &companyUpdateRequest)

	companyID := params.ByName("companyID")
	id, err := strconv.Atoi(companyID)
	helper.PanicIfError(err)

	companyUpdateRequest.ID = id

	companyResponse := controller.CompanyService.Update(r.Context(), companyUpdateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   companyResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CompanyControllerImpl) FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	companyID := params.ByName("companyID")
	ID, err := strconv.Atoi(companyID)
	helper.PanicIfError(err)

	companyResponse := controller.CompanyService.FindByID(r.Context(), ID)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   companyResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CompanyControllerImpl) FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	companyResponses := controller.CompanyService.FindAll(r.Context())
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   companyResponses,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
