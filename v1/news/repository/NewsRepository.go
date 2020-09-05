/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import "github.com/MbungeApp/mbunge-core/models/db"

type NewsRepository interface {
	GetAllNews() ([]db.EventNew, error)
}
