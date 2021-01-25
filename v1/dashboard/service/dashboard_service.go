package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type DashboardServices interface {
	GetMetrics() response.Metrics
	ViewAllEvents() ([]db.EventNew, error)
	ViewEventById(id string) db.EventNew
	AddEvent(event *request.EventRequest) error
	EditEvent(id string, event *request.EventRequest) error
	DeleteEvent(id string) error
}
