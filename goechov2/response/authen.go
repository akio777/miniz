package response

import "net/http"

func InvalidBody() (int, interface{}) {
	return http.StatusBadRequest, Response{Message: "invalid body"}
}
