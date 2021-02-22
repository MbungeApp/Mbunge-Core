/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	deviceDao dao.NewDeviceDaoInterface
	userDao   dao.NewUserDaoInterface
}

func NewUserRepository(client *mongo.Client) UserRepository {

	userDao := dao.NewUserDaoInterface{
		Client: client,
	}
	deviceDao := dao.NewDeviceDaoInterface{
		Client: client,
	}
	return &userRepository{
		userDao:   userDao,
		deviceDao: deviceDao,
	}

}

// Device implementations
func (u userRepository) AddDevice(device db.Device) (db.Device, error) {
	device, err := u.deviceDao.AddDevice(device)
	if err != nil {
		return db.Device{}, err
	}
	return device, nil
}

func (u userRepository) UpdateDevice(id string, key string, value string) (db.Device, error) {
	device, err := u.deviceDao.UpdateDevice(id, key, value)
	if err != nil {
		return db.Device{}, err
	}
	return device, nil
}

func (u userRepository) GetDevice(userId string) (db.Device, error) {
	device, err := u.deviceDao.GetDevice(userId)
	if err != nil {
		return db.Device{}, err
	}
	return device, err
}

// User Implementation
func (u userRepository) AddUser(user db.User) (db.User, error) {
	user, err := u.userDao.AddUser(user)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (u userRepository) GetOneUser(id string) (db.User, error) {
	user, err := u.userDao.GetUserById(id)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (u userRepository) GetOneUserByPhone(phone string) (db.User, error) {
	user, err := u.userDao.GetUserByPhone(phone)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (u userRepository) DeleteUser(phone string) error {
	err := u.userDao.DeleteUser(phone)
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) UpdateUserVerification(phone string) error {
	panic("implement me")
}

func (u userRepository) UpdateUser(id string, key string, value string) error {
	err := u.userDao.UpdateUser(id, key, value)
	if err != nil {
		return err
	}
	return nil
}
