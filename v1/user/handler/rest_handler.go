/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package handler

import (
	"log"
	"net/http"

	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/user/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ms "github.com/mitchellh/mapstructure"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(e *echo.Echo, userService service.UserService) {
	userRestHandler := &userRestHandler{userService: userService}
	g := e.Group("/api/v1/auth")
	g.POST("/sign_in", userRestHandler.SignInUser)
	g.POST("/sign_up", userRestHandler.SignUpUser)
	g.Use(middleware.JWT([]byte("secret")))
}

// @Summary Sign-in User
// @Description register user
// @Produce json
// @Accept json
// @Param phone body string true "254727751832", password body string true "1234567890"
// @Success 200 {object} object "ok" response.LoginResponse
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /api/v1/auth/sign_in [post]
func (u *userRestHandler) SignInUser(c echo.Context) error {
	var userRequest request.LoginRequest
	var info echo.Map
	var err error

	// 1. bind the request payload to a struct
	err = c.Bind(&info)
	decodeConfig := ms.DecoderConfig{TagName: "json", Result: &userRequest}
	decoder, err := ms.NewDecoder(&decodeConfig)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = decoder.Decode(&info)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// 2. Login the user
	res, err := u.userService.LoginUser(userRequest)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	// 3. Return the token and current logged in user info
	return c.JSON(http.StatusOK, res)
}

func (u *userRestHandler) SignUpUser(c echo.Context) error {
	var registerReq request.RegisterRequest
	var info echo.Map
	var err error

	//1. bind the request payload to a struct
	err = c.Bind(&info)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusConflict, err)
	}
	decodeConfig := ms.DecoderConfig{TagName: "json", Result: &registerReq}
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

	// 2 . send the data to service layer
	res, err := u.userService.RegisterUser(registerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}
