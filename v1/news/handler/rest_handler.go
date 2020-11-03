/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package handler

import (
	"net/http"

	"github.com/MbungeApp/mbunge-core/v1/news/service"
	"github.com/labstack/echo/v4"
)

type eventsRestHandler struct {
	eventService service.NewsService
}

func NewEventRestHandler(e *echo.Echo, eventService service.NewsService) {
	eventsRestHandler := &eventsRestHandler{eventService: eventService}
	g := e.Group("/api/v1/events")
	//g.Use(middleware.JWT([]byte("secret")))
	g.GET("/", eventsRestHandler.allEvents)
}

// Implemented methods
func (e *eventsRestHandler) allEvents(c echo.Context) error {
	events, err := e.eventService.AllNews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, events)
}
