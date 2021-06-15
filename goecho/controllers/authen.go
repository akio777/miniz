package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Authenticate struct{}

func (x *Authenticate) Login(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
