package main

import (
	"backend/controllers"
	"backend/database"
	"backend/functions"
	"backend/middlewares"
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
		Format:           "[${time_custom}]  Method : ${method}  Status : ${status}  URL : ${uri} \n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	app.Use(middleware.Recover())

	database.Init_table()

	app.GET("/", HealthCheck)
	app.GET("/check_test", Test)

	authen := app.Group("/authen")
	// authen.Use(middleware.BasicAuth(controllers.Auth.BasicAuth))
	authen.GET("/login", controllers.Auth.Login)
	authen.GET("/register", controllers.Auth.Register)

	users := app.Group("/services")
	// users.Use(middleware.JWT([]byte(functions.GoDotEnv("SECRET"))))
	users.Use(middlewares.CheckToken())
	users.GET("/", Test)
	// users.GET("", controllers.Auth.Test)

	// app.GET("/_test", controllers.Auth.Test)

	PORT := ":" + functions.GoDotEnv("PORT")
	app.Logger.Fatal(app.Start(PORT))
}
