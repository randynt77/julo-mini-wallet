package util

import (
	"encoding/json"
	"net/http"
)

type (
	RestStandard struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
	}
	ErrorResponse struct {
		Error string `json:"error,omitempty"`
	}
)

func NewRestResponse() (res RestStandard) {
	r := RestStandard{}
	return r
}

func (res RestStandard) Write(w http.ResponseWriter, httpStatus int) {
	byteData, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(byteData)
}

func (res RestStandard) WriteError(w http.ResponseWriter, httpStatus int, errorMessage string) {
	res.Status = "fail"
	errResp := ErrorResponse{
		Error: errorMessage,
	}
	res.Data = errResp
	res.Write(w, httpStatus)
}
