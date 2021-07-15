package svc

import (
	"backend/database"
	"backend/models"
	"backend/response"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	user := new(models.Users)
	if err := e.Bind(&user); err != nil {
		return e.JSON(response.InvalidBody())
	}
	if err := User.Create(database.DB, user); err != nil {
		return e.JSON(response.UserConflict())
	}
	return e.JSON(response.Success())
}

func (x *_Authen) Login(e echo.Context) error {
	username := e.FormValue("username")
	password := BytePassword(e.FormValue("password"))
	user, err := User.ReadByUsername(database.DB, username)
	if err != nil {
		return e.JSON(response.UserNotfound())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return e.JSON(response.InvalidPassword())
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	fmt.Println(claims["exp"])
	tk, err := token.SignedString([]byte(GoDotEnv("SECRET")))
	if err != nil {
		return e.JSON(response.InternalError("can not signed token"))
	}
	return e.JSON(response.OK("", tk))
}
