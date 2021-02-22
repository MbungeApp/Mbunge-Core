/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/models/db"
)

type UserRepository interface {
	AddDevice(device db.Device) (db.Device, error)
	UpdateDevice(id string, key string, value string) (db.Device, error)
	GetDevice(userId string) (db.Device, error)
	AddUser(user db.User) (db.User, error)
	GetOneUser(id string) (db.User, error)
	GetOneUserByPhone(phone string) (db.User, error)
	UpdateUser(id string, key string, value string) error
	UpdateUserVerification(phone string) error
	DeleteUser(phone string) error
}
