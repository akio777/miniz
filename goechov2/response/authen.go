package response

import (
	"net/http"
)

func InvalidBody() (int, interface{}) {
	return http.StatusBadRequest, Response{Message: "invalid body"}
}
func UserConflict() (int, interface{}) {
	return http.StatusConflict, Response{Message: "user already exists"}
}
