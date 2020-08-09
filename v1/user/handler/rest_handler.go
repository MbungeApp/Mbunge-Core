package handler

import (
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/user/service"
	"github.com/labstack/echo"
	ms "github.com/mitchellh/mapstructure"
	"log"
	"net/http"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(e *echo.Echo, userService service.UserService) {
	userRestHandler := &userRestHandler{userService: userService}
	g := e.Group("/api/v1/auth")
	g.POST("/sign_in", userRestHandler.SignInUser)
	g.POST("/sign_up", userRestHandler.SignUpUser)
}

func (u *userRestHandler) SignInUser(c echo.Context) error {
	var userRequest request.LoginRequest
	var info echo.Map
	var err error
	log.Println(info)

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
