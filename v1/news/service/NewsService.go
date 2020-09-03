package service

import "github.com/MbungeApp/mbunge-core/models/db"

type NewsService interface {
	AllNews() ([]db.EventNew, error)
}
