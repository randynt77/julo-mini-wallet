package util

import (
	"encoding/json"
	"net/http"
)

type (
	RestStandard struct {
		Status  string      `json:"status"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
		Code    int         `json:"code,omitempty"`
	}
	ErrorResponse struct {
		error string
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
	res.Data = ErrorResponse{
		error: errorMessage,
	}
	res.Write(w, httpStatus)
}
