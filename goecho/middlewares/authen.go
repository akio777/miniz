package middlewares

import (
	"github.com/labstack/echo/v4"
)

func CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			header := e.Request().Header.Get("Authorization")
			if header != "" {

			}
			// if header != "" {
			// 	user := e.Get("user").(*jwt.Token)
			// 	claims := user.Claims.(jwt.MapClaims)
			// 	timestamp := claims["exp"].(float64)
			// 	check := time.Unix(int64(timestamp), 0)
			// 	remainder := check.Sub(time.Now())

			// 	if !(remainder > 0) {
			// 		// return errors.WithMessage(echo.ErrBadRequest, "token expired, please re-login")
			// 		return e.JSON(http.StatusUnauthorized, models.Response{Message: "token not found"})
			// 	}
			// } else {
			// 	// return errors.WithMessage(echo.ErrBadRequest, "token not found")
			// 	return e.JSON(http.StatusUnauthorized, models.Response{Message: "token not found"})
			// }
			return next(e)
		}
	}
}
