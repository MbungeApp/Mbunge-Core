/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MbungeApp/mbunge-core/v1/participation/service"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type participationWSHandler struct {
	participationService service.ParticipationService
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewWebsocketHandler(e *echo.Echo, participationService service.ParticipationService) {
	participationWSHandler := participationWSHandler{
		participationService: participationService,
	}
	g := e.Group("/api/v1/participation")
	g.GET("/ws", participationWSHandler.ChangeStreams)
}

// Websocket handler
func (p *participationWSHandler) ChangeStreams(c echo.Context) error {
	if c.IsWebSocket() {

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		changes, err := p.participationService.ChangeStreams()
		for changes.Next(context.TODO()) {
			fmt.Println(changes.Current)
			err = ws.WriteMessage(websocket.TextMessage, changes.Current)
			if err != nil {
				c.Logger().Error(err)
			}
		}

	} else {
		return c.JSON(http.StatusBadRequest, "websocket channel")
	}
	return nil
}
