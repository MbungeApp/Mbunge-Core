package main

import (
	"github.com/MbungeApp/mbunge-core/config"
	userHandler "github.com/MbungeApp/mbunge-core/v1/user/handler"
	userRepo "github.com/MbungeApp/mbunge-core/v1/user/repository"
	userService "github.com/MbungeApp/mbunge-core/v1/user/service"

	eventHandler "github.com/MbungeApp/mbunge-core/v1/news/handler"
	eventRepo "github.com/MbungeApp/mbunge-core/v1/news/repository"
	eventService "github.com/MbungeApp/mbunge-core/v1/news/service"

	participationHandler "github.com/MbungeApp/mbunge-core/v1/participation/handler"
	participationRepo "github.com/MbungeApp/mbunge-core/v1/participation/repository"
	participationService "github.com/MbungeApp/mbunge-core/v1/participation/service"

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

	// User
	userRepository := userRepo.NewUserRepository(client)
	userservice := userService.NewUserServiceImpl(userRepository)
	userHandler.NewUserRestHandler(e, userservice)

	// Events
	eventRepository := eventRepo.NewEventRepository(client)
	eventservice := eventService.NewEventService(eventRepository)
	eventHandler.NewEventRestHandler(e, eventservice)

	// Participation
	participationRepository := participationRepo.NewParticipationRepositoryImpl(client)
	participationservice := participationService.NewParticipationServiceImpl(participationRepository)
	participationHandler.NewParticipationRestHandler(e, participationservice)
	participationHandler.NewWebsocketHandler(e, participationservice)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":" + "5000"))
}
