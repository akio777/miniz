package response

import (
	"net/http"
)

func UserConflict() (int, interface{}) {
	return http.StatusConflict, Response{Message: "user already exists"}
}
func UserNotfound() (int, interface{}) {
	return http.StatusNotFound, Response{Message: "user not exists"}
}
func InvalidPassword() (int, interface{}) {
	return http.StatusUnauthorized, Response{Message: "password invalid"}
}
