package request

import "github.com/MbungeApp/mbunge-core/models/db"

type RegisterRequest struct {
	Device db.Device
	User   db.User
}
