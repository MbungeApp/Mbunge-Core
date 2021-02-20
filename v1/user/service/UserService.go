/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type UserService interface {
	RegisterUser(regRequest *request.RegisterRequest) (response.RegisterResponse, error)
	LoginUser(request request.LoginRequest) (response.LoginResponse, error)
	EditUser(id string, user *request.EditUser) error
	ViewUserById(id string) db.User
	DeleteUser(id string) error
	VerifyAccount(phone string) error
	PasswordReset(phone string) error
}
