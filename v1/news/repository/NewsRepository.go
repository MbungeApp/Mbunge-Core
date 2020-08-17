package repository

import "github.com/MbungeApp/mbunge-core/models/db"

type NewsRepository interface {
	GetAllNews() ([]db.EventNew, error)
}
