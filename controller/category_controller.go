package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
}
