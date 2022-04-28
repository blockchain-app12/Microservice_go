package utils

import (
	"strings"
	"net/http"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(w http.ResponseWriter,statusCode int, data interface{}) Response {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := Response{
		Status:  true,
		Message: "API Success",
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(w http.ResponseWriter,statusCode int,message string, err string, ) Response {
	splittedError := strings.Split(err, "\n")
	w.WriteHeader(statusCode)
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    EmptyObj{},
	}
	return res
}
