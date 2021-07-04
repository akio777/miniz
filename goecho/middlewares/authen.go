package middlewares

import (
	// "backend/models"
	// "net/http"
	"fmt"
	// "time"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		fmt.Println("XXX")
		// header := e.Request().Header.Get("Authorization")
		// if header != "" {
		// 	user := e.Get("user").(*jwt.Token)
		// 	claims := user.Claims.(jwt.MapClaims)
		// 	timestamp := claims["exp"].(float64)
		// 	check := time.Unix(int64(timestamp), 0)
		// 	remainder := check.Sub(time.Now())
		// 	if !(remainder > 0) {
		// 		return errors.WithMessage(echo.ErrBadRequest, "token expired, please re-login")
		// 	}
		// } else {
		// 	return errors.WithMessage(echo.ErrBadRequest, "token not found")
		// 	// return e.JSON(http.StatusUnauthorized, models.Response{Message: "token not found"})
		// }
		return next(e)
	}
}
