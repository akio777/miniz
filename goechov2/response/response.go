package response

import "net/http"

type Response struct {
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

func Success() (int, interface{}) {
	return http.StatusCreated, Response{Message: "create success"}
}

func OK(msg string, data interface{}) (int, interface{}) {
	return http.StatusOK, Response{Message: msg, Data: data}
}
func InvalidBody() (int, interface{}) {
	return http.StatusBadRequest, Response{Message: "invalid body"}
}
func InternalError(msg string) (int, interface{}) {
	return http.StatusInternalServerError, Response{Message: msg}
}
