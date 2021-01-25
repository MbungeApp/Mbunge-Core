package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/dashboard/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type dashboardHandler struct {
	dashboardService service.DashboardServices
}

func NewDashboardHandler(e *echo.Echo, dashService service.DashboardServices) {
	dashboardHandler := &dashboardHandler{
		dashboardService: dashService,
	}
	g := e.Group("/api/v1/dashboard")
	g.GET("/metrics", dashboardHandler.home)
	g.GET("/event/", dashboardHandler.getAllEvents)
	g.POST("/event/add", dashboardHandler.addEvent)
	g.GET("/event/view/:id", dashboardHandler.getOneEvent)
	g.POST("/event/edit/:id", dashboardHandler.editEvent)
	g.DELETE("/event/delete/:id", dashboardHandler.deleteEvent)
}

// Dashboard
// 1. Metrics
func (d dashboardHandler) home(c echo.Context) error {
	var metrics interface{}
	marshalMetrics, err := json.Marshal(d.dashboardService.GetMetrics())
	if err != nil {
		fmt.Println("error: " + err.Error())
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err = json.Unmarshal(marshalMetrics, &metrics)
	return c.JSON(http.StatusOK, metrics)
}

// Events
// 1. Get All
// 2. Add
// 3. View one
// 4. Edit
// 5. Delete
func (d dashboardHandler) getAllEvents(c echo.Context) error {
	var eventMap []interface{}
	eventsDb, err := d.dashboardService.ViewAllEvents()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	marshalEvents, err := json.Marshal(eventsDb)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err = json.Unmarshal(marshalEvents, &eventMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	return c.JSON(http.StatusOK, eventMap)

}
func (d dashboardHandler) addEvent(c echo.Context) error {
	eventReq := new(request.EventRequest)
	if err := c.Bind(eventReq); err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.AddEvent(eventReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	return c.JSON(http.StatusFound, "Added successfully")
}
func (d dashboardHandler) getOneEvent(c echo.Context) error {
	id := c.Param("id")
	event := d.dashboardService.ViewEventById(id)
	return c.JSON(http.StatusOK, event)
}
func (d dashboardHandler) editEvent(c echo.Context) error {
	id := c.Param("id")

	var eventMap interface{}
	eventReq := new(request.EventRequest)
	if err := c.Bind(eventReq); err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.EditEvent(id, eventReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	even := d.dashboardService.ViewEventById(id)

	marshalParticipations, err := json.Marshal(even)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err = json.Unmarshal(marshalParticipations, &eventMap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}

	return c.JSON(http.StatusOK, eventMap)
}

func (d dashboardHandler) deleteEvent(c echo.Context) error {
	id := c.Param("id")

	err := d.dashboardService.DeleteEvent(id)
	if err != nil {
		fmt.Println("Error deleting: ", err.Error())
		return c.String(http.StatusInternalServerError, "error occurred")
	}

	return c.JSON(http.StatusFound, "deleted successfully")
}
