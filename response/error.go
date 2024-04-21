package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json: "message"`
}

func InternalServerError(msg string) ErrorResponse {
	return error(msg, http.StatusInternalServerError)
}

func BadRequest(msg string) ErrorResponse {
	return error(msg, http.StatusBadRequest)
}

func NotFound(msg string) ErrorResponse {
	return error(msg, http.StatusNotFound)
}

func Unauthorized(msg string) ErrorResponse {
	return error(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) ErrorResponse {
	return error(msg, http.StatusForbidden)
}

func error(msg string, code int) ErrorResponse {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ErrorResponse) GetData() interface{} {
	return nil
}
