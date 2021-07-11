package svc

import (
	"backend/database"
	"backend/models"
	"backend/response"
	"net/http"

	// "golang.org/x/crypto/bcrypt"
	// "time"

	"github.com/labstack/echo/v4"
)

type _Authen struct {
}

func (x *_Authen) BasicAuth(username string, password string, e echo.Context) (bool, error) {
	if username == "dev" && password == "123123" {
		return true, nil
	}
	return false, echo.ErrUnauthorized
}

func (x *_Authen) Register(e echo.Context) error {

	user := new(models.User_info)
	if err := e.Bind(&user); err != nil {
		return e.JSON(response.InvalidBody())
	}
	if err := User.Create(database.DB, user); err != nil {
		return e.JSON(response.UserConflict())
	}
	return e.JSON(response.Success())
}

func (x *_Authen) Login(e echo.Context) error {
	// password := svc.BytePassword(e.FormValue("password"))
	// err := bcrypt.CompareHashAndPassword(, password)
	// if err != nil {
	// 	return e.JSON(http.StatusNotFound, "")
	// }
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["user_id"] = 1
	// claims["admin"] = true
	// claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	// fmt.Println(claims["exp"])
	// tk, err := token.SignedString([]byte(svc.GoDotEnv("SECRET")))
	// if err != nil {
	// 	panic(err)
	// }
	return e.JSON(http.StatusOK, "")
}
