package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Save(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByID(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
}
