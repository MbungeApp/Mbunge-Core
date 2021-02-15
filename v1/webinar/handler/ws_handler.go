/*
* Copyright (c) 2020.
* MbungeApp Inc all rights reserved
 */

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/MbungeApp/mbunge-core/v1/webinar/service"
	"github.com/gorilla/websocket"
)

type webinarWSHandler struct {
	webinarService service.WebinarService
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewWebsocketHandler(e *echo.Echo, webinarService service.WebinarService) {
	webinarwsHandler := webinarWSHandler{
		webinarService: webinarService,
	}
	g := e.Group("/api/v1/webinar")
	g.GET("/ws", webinarwsHandler.ChangeStreams)
}

// Websocket handler
func (p *webinarWSHandler) ChangeStreams(c echo.Context) error {
	if c.IsWebSocket() {

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		changes, err := p.webinarService.ChangeStreams()
		fmt.Println("waiting for changes")
		var changeDoc map[string]interface{}
		for changes.Next(context.Background()) {
			if e := changes.Decode(&changeDoc); e != nil {
				fmt.Printf("error decoding: %s\n", e)
			}
			fmt.Println("change: ")
			webinars := p.webinarService.GetAllWebinars()
			bytes, err := json.Marshal(webinars)
			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				c.Logger().Error(err)
			}
		}

	} else {
		return c.JSON(http.StatusBadRequest, "websocket channel")
	}
	return nil
}
