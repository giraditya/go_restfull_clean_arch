package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(rw http.ResponseWriter, response interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(response)
	PanicIfError(err)
}
