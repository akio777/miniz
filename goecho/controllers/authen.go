package controllers

import (
	"backend/functions"
	"backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Authenticate struct{}

func (x *Authenticate) BasicAuth(username string, password string, e echo.Context) (bool, error) {
	if username == "dev" && password == "123123" {
		return true, nil
	}
	return false, echo.ErrUnauthorized
}

var TempHash []byte

func BytePassword(password string) []byte {
	return []byte(password)
}

func (x *Authenticate) Register(e echo.Context) error {
	// username := e.FormValue("username")
	// TODO check this username in DB if exists return
	// return e.JSON(http.StatusConflict, models.Response{Message: "user already exists"})
	password := BytePassword(e.FormValue("password"))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	// TODO insert Users password with this hash
	TempHash = hashedPassword
	return e.JSON(http.StatusCreated, models.Response{Message: "user created"})
}

func (x *Authenticate) Login(e echo.Context) error {
	// username := e.FormValue("username")
	// TODO query Users.password table where username = username
	// * if not exists
	// * return e.JSON(http.StatusNotFound, models.Response{Message: "invalid username or password, please check and login again"})
	// !TempHash must be password from Users.password
	password := BytePassword(e.FormValue("password"))
	err := bcrypt.CompareHashAndPassword(TempHash, password)
	if err != nil {
		return e.JSON(http.StatusNotFound, models.Response{Message: "invalid username or password, please check and login again"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 1
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	fmt.Println(claims["exp"])
	tk, err := token.SignedString([]byte(functions.GoDotEnv("SECRET")))
	if err != nil {
		panic(err)
	}
	return e.JSON(http.StatusOK, models.Response{Message: "login success", Data: tk})
}

func (x *Authenticate) Check(e echo.Context) error {
	header := e.Request().Header.Get("Authorization")
	if header != "" {
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		timestamp := claims["exp"].(float64)
		check := time.Unix(int64(timestamp), 0)
		remainder := check.Sub(time.Now())
		if !(remainder > 0) {
			return e.JSON(http.StatusUnauthorized, models.Response{Message: "token expired, please re-login"})
		}

		return e.JSON(http.StatusOK, models.Response{})
	} else {
		return e.JSON(http.StatusUnauthorized, models.Response{Message: "token not found"})
	}
}
