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
	GetOneUser(phone string) (db.User, error)
	UpdateUser(id string, key string, value string) (db.Device, error)
	UpdateUserVerification(phone string) error
}
