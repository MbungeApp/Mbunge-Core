/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/utils"
	"github.com/MbungeApp/mbunge-core/v1/user/repository"
	"log"
	"math/rand"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: repository,
	}
}

func (u userServiceImpl) RegisterUser(regRequest request.RegisterRequest) (response.RegisterResponse, error) {
	var device db.Device
	user := regRequest.User

	userres, err := u.userRepository.AddUser(user)
	if err != nil {
		return response.RegisterResponse{}, err
	}
	go func() {
		device.UserId = userres.ID.Hex()
		device.Type = regRequest.Type
		device.FCMToken = regRequest.FcmToken
		_, err = u.userRepository.AddDevice(device)
		if err != nil {
			log.Fatal(err)
		}
	}()

	res := response.RegisterResponse{
		Code: rand.Intn(9999),
		User: userres,
	}
	return res, nil
}

func (u userServiceImpl) LoginUser(request request.LoginRequest) (response.LoginResponse, error) {
	phone := request.Phone
	password := request.Password

	user, err := u.userRepository.GetOneUser(phone)
	if err != nil {
		return response.LoginResponse{}, err
	}
	match, err := utils.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return response.LoginResponse{}, err
	}
	if match {
		// new token
		token, err := utils.GenerateToken(user.EmailAddress)
		if err != nil {
			return response.LoginResponse{}, err
		}
		res := response.LoginResponse{
			Token: token,
			User:  user,
		}
		return res, nil
	} else {
		return response.LoginResponse{}, nil
	}
}

func (u userServiceImpl) VerifyAccount(phone string) error {
	panic("implement me" + phone)
}

func (u userServiceImpl) PasswordReset(phone string) error {
	panic("implement me" + phone)
}
