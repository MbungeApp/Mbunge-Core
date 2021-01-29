/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package main

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/config"
	userHandler "github.com/MbungeApp/mbunge-core/v1/user/handler"
	userRepo "github.com/MbungeApp/mbunge-core/v1/user/repository"
	_userService "github.com/MbungeApp/mbunge-core/v1/user/service"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"net"

	eventHandler "github.com/MbungeApp/mbunge-core/v1/news/handler"
	eventRepo "github.com/MbungeApp/mbunge-core/v1/news/repository"
	_eventService "github.com/MbungeApp/mbunge-core/v1/news/service"

	participationHandler "github.com/MbungeApp/mbunge-core/v1/webinar/handler"
	participationRepo "github.com/MbungeApp/mbunge-core/v1/webinar/repository"
	_participationService "github.com/MbungeApp/mbunge-core/v1/webinar/service"

	mpHandler "github.com/MbungeApp/mbunge-core/v1/mp/handler"
	mpRepo "github.com/MbungeApp/mbunge-core/v1/mp/repository"
	_mpService "github.com/MbungeApp/mbunge-core/v1/mp/service"

	_dashboardHandler "github.com/MbungeApp/mbunge-core/v1/dashboard/handler"
	_dashboardService "github.com/MbungeApp/mbunge-core/v1/dashboard/service"

	"net/http"

	_ "github.com/MbungeApp/mbunge-core/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	e        *echo.Echo
	mqttConn mqtt.Client
	client   *mongo.Client
)

func init() {
	e = echo.New()
	mqttConn = config.ConnectMqtt()
	client = config.ConnectDB()
}

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

	// middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// User
	userRepository := userRepo.NewUserRepository(client)
	userService := _userService.NewUserServiceImpl(userRepository)
	userHandler.NewUserRestHandler(e, userService)

	// Events
	eventRepository := eventRepo.NewEventRepository(client)
	eventService := _eventService.NewEventService(eventRepository)
	eventHandler.NewEventRestHandler(e, eventService)

	// Participation
	participationRepository := participationRepo.NewParticipationRepositoryImpl(client)
	participationService := _participationService.NewParticipationServiceImpl(participationRepository)
	participationHandler.NewParticipationRestHandler(e, participationService)
	participationHandler.NewWebsocketHandler(e, participationService)
	webinar := participationHandler.NewMqttWebinarHandler(&mqttConn, participationService)
	// MP
	mpRepository := mpRepo.NewMpRepository(client)
	mpService := _mpService.NewMpService(mpRepository)
	mpHandler.NewMpRestHandler(e, mpService)

	// Dashboard
	dashboardService := _dashboardService.NewDashboardServiceImpl(client)
	_dashboardHandler.NewDashboardHandler(e, dashboardService)

	// Swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// mqtt
	go listen(mqttConn, "topics/server", webinar.SystemActions)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	l, err := net.Listen("tcp4", ":5000")
	if err != nil {
		e.Logger.Fatal(l)
	}
	e.Listener = l
	e.Logger.Fatal(e.Start(""))
}

func listen(client mqtt.Client, topic string, handler mqtt.MessageHandler) {
	token := client.Subscribe(topic, 0, handler)
	if token != nil {
		fmt.Printf("Subscribed to topic: %s\n", topic)
	} else {
		fmt.Printf("Failed to subscribe to topic: %s\n", topic)
	}
}
