package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/v1/user/repository"
	"log"
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

	response := response.RegisterResponse{
		Code: 1111,
		User: userres,
	}
	return response, nil
}

func (u userServiceImpl) LoginUser(request request.LoginRequest) (response.LoginResponse, error) {
	phone := request.Phone
	password := request.Password

	user, err := u.userRepository.GetOneUser(phone)
	if err != nil {
		return response.LoginResponse{}, err
	}
	if user.Password == password {
		response := response.LoginResponse{
			Token: "1234567", //TODO
			User:  user,
		}
		return response, nil
	} else {
		return response.LoginResponse{}, nil
	}
}

func (u userServiceImpl) VerifyAccount(phone string) error {
	panic("implement me")
}

func (u userServiceImpl) PasswordReset(phone string) error {
	panic("implement me")
}
