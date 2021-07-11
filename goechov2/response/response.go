package response

import "net/http"

type Response struct {
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

func Success() (int, interface{}) {
	return http.StatusCreated, Response{Message: "create success"}
}
