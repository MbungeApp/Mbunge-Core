package main

import (
	"github.com/MbungeApp/mbunge-core/config"
	"github.com/MbungeApp/mbunge-core/v1/user/handler"
	"github.com/MbungeApp/mbunge-core/v1/user/repository"
	"github.com/MbungeApp/mbunge-core/v1/user/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	client := config.ConnectDB()

	// middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	userRepository := repository.NewUserRepository(client)
	userService := service.NewUserServiceImpl(userRepository)
	handler.NewUserRestHandler(e, userService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":" + "5000"))
}
