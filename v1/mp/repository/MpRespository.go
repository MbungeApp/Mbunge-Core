/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type MpRepository interface {
	GetAllMps() ([]db.MP, error)
	GetMpOftheWeek() response.MpOftheWeek
}
