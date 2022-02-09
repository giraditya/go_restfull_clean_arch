package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	RequestToken(rw http.ResponseWriter, r *http.Request, params httprouter.Params)
}
