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

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	produuctCreateRequest := request.ProductCreateRequest{}
	helper.ReadFromRequestBody(r, &produuctCreateRequest)

	productResponse := controller.ProductService.Save(r.Context(), produuctCreateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *ProductControllerImpl) Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productUpdateRequest := request.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	productID := params.ByName("productID")
	id, err := strconv.Atoi(productID)
	helper.PanicIfError(err)

	productUpdateRequest.ID = id

	productResponse := controller.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *ProductControllerImpl) Delete(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productID := params.ByName("productID")
	id, err := strconv.Atoi(productID)
	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *ProductControllerImpl) FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productID := params.ByName("productID")
	id, err := strconv.Atoi(productID)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindByID(r.Context(), id)
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *ProductControllerImpl) FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.FindAll(r.Context())
	webResponse := api.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
