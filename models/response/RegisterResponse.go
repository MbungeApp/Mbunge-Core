package response

import "github.com/MbungeApp/mbunge-core/models/db"

type RegisterResponse struct {
	Code int
	User db.User
}
