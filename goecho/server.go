package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"Pea-backend/controllers"
	"Pea-backend/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Server is up and running",
	})
}

func Test(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "nothing",
	})
}

func main() {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ::: ${uri} ::: ${method} ::: status - ${status} \n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	app.Use(middleware.Recover())

	database.Init_table()
	database.Init_data()

	app.GET("/", HealthCheck)
	app.GET("/test", Test)
	authen := app.Group("/auth")
	authen.GET("/login", controllers.Auth.Login)

	app.Logger.Fatal(app.Start(":7001"))
}
