package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type MpService interface {
	AllMps() ([]db.MP, error)
	MpOftheWeek() response.MpOftheWeek
}
