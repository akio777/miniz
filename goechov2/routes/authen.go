package routes

import (
	"backend/svc"

	"github.com/labstack/echo/v4"
)

func AuthenRoute(g *echo.Group) {
	g.GET("/login", svc.Authen.Login)
	g.POST("/register", svc.Authen.Register)
}
