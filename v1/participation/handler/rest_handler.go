/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package handler

import (
	"log"
	"net/http"

	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/participation/service"
	"github.com/labstack/echo/v4"
	ms "github.com/mitchellh/mapstructure"
)

type participationRestHandler struct {
	participationService service.ParticipationService
}

func NewParticipationRestHandler(e *echo.Echo, participationService service.ParticipationService) {
	participationRestHandler := participationRestHandler{participationService: participationService}

	// semantic versioning of api !!
	g := e.Group("/api/v1/participation")
	//g.Use(middleware.JWT([]byte("secret")))
	g.GET("/", participationRestHandler.AllParticipation)
	g.GET("/response/:id", participationRestHandler.AllResponseByParticipation)
	g.POST("/response/add", participationRestHandler.AddResponse)
	g.DELETE("/response/delete/:id", participationRestHandler.DeleteResponse)
}

// Rest handlers

// @Summary All Partipations
// @Description get all
// @Produce json
// @Accept json
// @Success 200 {array} array "ok" []db.Participation
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /api/v1/participation/ [get]
func (p *participationRestHandler) AllParticipation(c echo.Context) error {
	return c.JSON(http.StatusOK, p.participationService.GetAllParticipation())
}
func (p *participationRestHandler) AllResponseByParticipation(c echo.Context) error {
	participationID := c.Param("id")
	return c.JSON(http.StatusOK, p.participationService.GetAllResponseByParti(participationID))
}
func (p *participationRestHandler) AddResponse(c echo.Context) error {
	var responseReq request.ResponseRequest
	var info echo.Map
	var err error

	//1. bind the request payload to a struct
	err = c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}
	decodeConfig := ms.DecoderConfig{TagName: "json", Result: &responseReq}
	decoder, err := ms.NewDecoder(&decodeConfig)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	err = decoder.Decode(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = p.participationService.AddResponse(responseReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "")
}
func (p *participationRestHandler) DeleteResponse(c echo.Context) error {
	responseID := c.Param("response_id")
	return c.JSON(http.StatusOK, p.participationService.DeleteResponse(responseID))
}
