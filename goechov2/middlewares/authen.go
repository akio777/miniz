package middlewares

import "github.com/labstack/echo/v4"

type _Authenticate struct{}

func (x *_Authenticate) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
