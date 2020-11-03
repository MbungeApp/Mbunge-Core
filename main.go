/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package main

import (
	"net"

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

	mpHandler "github.com/MbungeApp/mbunge-core/v1/mp/handler"
	mpRepo "github.com/MbungeApp/mbunge-core/v1/mp/repository"
	mpService "github.com/MbungeApp/mbunge-core/v1/mp/service"

	"net/http"

	_ "github.com/MbungeApp/mbunge-core/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Mbunge App API
// @version 1.0
// @description Version one of the api.
// @termsOfService http://dashboard.mbungeapp.tech/terms/

// @contact.name API Support
// @contact.url http://dashboard.mbungeapp.tech/support
// @contact.email 858wpwaweru@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
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

	// MP
	mpRepository := mpRepo.NewMpRepository(client)
	mpservice := mpService.NewMpService(mpRepository)
	mpHandler.NewMpRestHandler(e, mpservice)

	// Swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	l, err := net.Listen("tcp4", ":5000")
	if err != nil {
		e.Logger.Fatal(l)
	}
	e.Listener = l
	e.Logger.Fatal(e.Start(""))
}
