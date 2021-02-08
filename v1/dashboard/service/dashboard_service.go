package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type DashboardServices interface {
	GetMetrics() response.Metrics

	// Webinar
	ViewAllWebinars() ([]db.Webinar, error)
	AddWebinar(webinar *request.AddWebinar) error
	DeleteWebinar(id string) error

	// Events
	ViewAllEvents() ([]db.EventNew, error)
	ViewEventById(id string) db.EventNew
	AddEvent(event *request.EventRequest) error
	EditEvent(id string, event *request.EventRequest) error
	DeleteEvent(id string) error

	// MPs
	ViewAllMps() ([]db.MP, error)
	ViewMpById(id string) db.MP
	AddMp(mp *request.MpRequest) error
	EditMp(id string, mp *request.MpRequest) error
	DeleteMp(id string) error

	// Management
	FetchAllAdmins() []db.Management
	FetchAdminById(id string) db.Management
	RegisterAdmin(regRequest request.AddManager) error
	LoginAdmin(request request.AdminLoginRequest) (response.AdminLoginResponse, error)
	EditAdmin(id string, admin *request.AddManager) error
	UpdateAdminPassword(password *request.UpdatePassword) (db.Management, error)
	DeleteAdmin(id string) error
}
