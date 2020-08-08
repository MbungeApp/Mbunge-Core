package response

import "github.com/MbungeApp/mbunge-core/models/db"

type LoginResponse struct {
	Token string
	User  db.User
}
