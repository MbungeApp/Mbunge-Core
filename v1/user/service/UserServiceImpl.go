/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/utils"
	"github.com/MbungeApp/mbunge-core/v1/user/repository"
	"log"
	"strconv"
	"time"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: repository,
	}
}

func (u userServiceImpl) RegisterUser(regRequest *request.RegisterRequest) (response.RegisterResponse, error) {
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
		Code: utils.GenerateOtp(),
		User: userres,
	}
	return res, nil
}

func (u userServiceImpl) LoginUser(request request.LoginRequest) (response.LoginResponse, error) {
	phone := request.Phone
	password := request.Password

	user, err := u.userRepository.GetOneUserByPhone(phone)
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

func (u userServiceImpl) DeleteUser(id string) error {
	err := u.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (u userServiceImpl) EditUser(id string, user *request.EditUser) error {
	originalUser, _ := u.userRepository.GetOneUser(id)
	fmt.Println(originalUser)

	if user.PhoneNumber != originalUser.PhoneNumber {
		err := u.userRepository.UpdateUser(id, "phone_number", user.PhoneNumber)
		if err != nil {
			return err
		}
	} else if user.EmailAddress != originalUser.EmailAddress {
		err := u.userRepository.UpdateUser(id, "email_address", user.EmailAddress)
		if err != nil {
			return err
		}
	} else if user.FirstName != originalUser.
		FirstName {
		err := u.userRepository.UpdateUser(id, "first_name", user.FirstName)
		if err != nil {
			return err
		}
	} else if user.LastName != originalUser.LastName {
		err := u.userRepository.UpdateUser(id, "last_name", user.LastName)
		if err != nil {
			return err
		}
	} else if user.County != originalUser.County {
		err := u.userRepository.UpdateUser(id, "county", user.County)
		if err != nil {
			return err
		}
	} else if user.DateBirth != originalUser.DateBirth {
		err := u.userRepository.UpdateUser(id, "date_birth", user.DateBirth.Format(time.RFC3339))
		if err != nil {
			return err
		}
	} else if user.Gender != originalUser.Gender {
		s := strconv.Itoa(user.Gender)
		err := u.userRepository.UpdateUser(id, "gender", s)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("******************nothing*******")
	}

	return nil
}

func (u userServiceImpl) ViewUserById(id string) db.User {
	user, err := u.userRepository.GetOneUser(id)
	if err != nil {
		return db.User{}
	}
	return user
}

func (u userServiceImpl) VerifyAccount(phone string) error {
	panic("implement me" + phone)
}

func (u userServiceImpl) PasswordReset(phone string) error {
	panic("implement me" + phone)
}
