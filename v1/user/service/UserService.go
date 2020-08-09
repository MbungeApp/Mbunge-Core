package service

import (
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type UserService interface {
	RegisterUser(regRequest request.RegisterRequest) (response.RegisterResponse, error)
	LoginUser(request request.LoginRequest) (response.LoginResponse, error)
	VerifyAccount(phone string) error
	PasswordReset(phone string) error
}
