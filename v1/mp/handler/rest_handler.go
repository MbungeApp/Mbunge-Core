/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package handler

import (
	"github.com/MbungeApp/mbunge-core/v1/mp/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type mpRestHandler struct {
	mpService service.MpService
}

func NewMpRestHandler(e *echo.Echo, mpService service.MpService) {
	mpRestHandler := &mpRestHandler{mpService: mpService}
	g := e.Group("/api/v1/mp")
	g.Use(middleware.JWT([]byte("secret")))
	g.GET("/", mpRestHandler.mpOftheWeek)
	g.GET("/all", mpRestHandler.allMps)
}

// Implemented methods
func (m *mpRestHandler) mpOftheWeek(c echo.Context) error {
	mps := m.mpService.MpOftheWeek()
	return c.JSON(http.StatusOK, mps)
}

func (m *mpRestHandler) allMps(c echo.Context) error {
	mps, err := m.mpService.AllMps()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, mps)
}
