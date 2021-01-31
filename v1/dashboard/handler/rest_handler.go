package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/dashboard/service"
	"github.com/labstack/echo/v4"
	ms "github.com/mitchellh/mapstructure"
	"log"
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
	// webinar
	g.GET("/webinar/", dashboardHandler.getAllWebinars)
	g.POST("/webinar/add", dashboardHandler.addWebinar)
	g.DELETE("/webinar/delete/:id", dashboardHandler.deleteWebinar)
	// Event
	g.GET("/event/", dashboardHandler.getAllEvents)
	g.POST("/event/add", dashboardHandler.addEvent)
	g.GET("/event/view/:id", dashboardHandler.getOneEvent)
	g.POST("/event/edit/:id", dashboardHandler.editEvent)
	g.DELETE("/event/delete/:id", dashboardHandler.deleteEvent)
	// Mp
	g.GET("/mp/", dashboardHandler.getAllMps)
	g.POST("/mp/add", dashboardHandler.addMp)
	g.GET("/mp/view/:id", dashboardHandler.getOneMp)
	g.POST("/mp/edit/:id", dashboardHandler.editMp)
	g.DELETE("/mp/delete/:id", dashboardHandler.deleteMp)
	// Management
	g.GET("/management/", dashboardHandler.getAllAdmins)
	g.POST("/management/sign_in", dashboardHandler.adminSignIn)
	g.POST("/management/add", dashboardHandler.addAdmin)
	g.POST("/management/edit/:id", dashboardHandler.editAdmin)
	g.DELETE("/management/delete/:id", dashboardHandler.deleteAdmin)
}

// ********************************
// Dashboard
// ********************************
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

// ********************************
// webinar
// ********************************
func (d dashboardHandler) getAllWebinars(c echo.Context) error {
	webinars, err := d.dashboardService.ViewAllWebinars()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, webinars)
}
func (d dashboardHandler) addWebinar(c echo.Context) error {
	webinarReq := new(request.AddWebinar)
	if err := c.Bind(webinarReq); err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.AddWebinar(webinarReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Added successfully")
}
func (d dashboardHandler) deleteWebinar(c echo.Context) error {
	webinarId := c.Param("id")
	err := d.dashboardService.DeleteWebinar(webinarId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}

// ********************************
// Events
// ********************************
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
	return c.JSON(http.StatusOK, "Added successfully")
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

	return c.JSON(http.StatusOK, "deleted successfully")
}

// ********************************
// MP
// ********************************
func (d dashboardHandler) getAllMps(c echo.Context) error {
	var mpInterface []interface{}
	mpsDb, err := d.dashboardService.ViewAllMps()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	marshalMps, err := json.Marshal(mpsDb)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err = json.Unmarshal(marshalMps, &mpInterface)
	return c.JSON(http.StatusOK, mpInterface)

}
func (d dashboardHandler) addMp(c echo.Context) error {
	mpReq := new(request.MpRequest)
	if err := c.Bind(mpReq); err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.AddMp(mpReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	return c.JSON(http.StatusOK, "added successfully")
}
func (d dashboardHandler) getOneMp(c echo.Context) error {
	return nil
}
func (d dashboardHandler) editMp(c echo.Context) error {
	id := c.Param("id")
	var mpMap interface{}
	mpReq := new(request.MpRequest)
	if err := c.Bind(mpReq); err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.EditMp(id, mpReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	mp := d.dashboardService.ViewMpById(id)

	marshalMp, err := json.Marshal(mp)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err = json.Unmarshal(marshalMp, &mpMap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error occurred")
	}

	return c.JSON(http.StatusOK, mpMap)
}
func (d dashboardHandler) deleteMp(c echo.Context) error {
	id := c.Param("id")

	err := d.dashboardService.DeleteMp(id)
	if err != nil {
		fmt.Println("Error deleting: ", err.Error())
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}

// ********************************
// Management
// ********************************
func (d dashboardHandler) getAllAdmins(c echo.Context) error {
	res := d.dashboardService.FetchAllAdmins()
	return c.JSON(http.StatusCreated, res)
}
func (d dashboardHandler) adminSignIn(c echo.Context) error {
	var adminLoginRequest request.AdminLoginRequest
	var info echo.Map
	var err error

	err = c.Bind(&info)
	decodeConfig := ms.DecoderConfig{TagName: "json", Result: &adminLoginRequest}
	decoder, err := ms.NewDecoder(&decodeConfig)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = decoder.Decode(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := d.dashboardService.LoginAdmin(adminLoginRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}
func (d dashboardHandler) addAdmin(c echo.Context) error {
	var adminAddReq request.AddManager
	var info echo.Map
	var err error

	//1. bind the request payload to a struct
	err = c.Bind(&info)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusConflict, err)
	}
	decodeConfig := ms.DecoderConfig{TagName: "json", Result: &adminAddReq}
	decoder, err := ms.NewDecoder(&decodeConfig)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	err = decoder.Decode(&info)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err = d.dashboardService.RegisterAdmin(adminAddReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "Added admin successfully")
}
func (d dashboardHandler) editAdmin(c echo.Context) error {
	id := c.Param("id")
	//var adminMap interface{}
	adminReq := new(request.AddManager)
	if err := c.Bind(adminReq); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "error occurred")
	}
	err := d.dashboardService.EditAdmin(id, adminReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	originalAdmin := d.dashboardService.FetchAdminById(id)
	return c.JSON(http.StatusOK, originalAdmin)
}
func (d dashboardHandler) deleteAdmin(c echo.Context) error {
	id := c.Param("id")

	err := d.dashboardService.DeleteAdmin(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "deleted successfully")
}
