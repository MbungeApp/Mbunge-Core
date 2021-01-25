package response

import "github.com/MbungeApp/mbunge-core/models/db"

type AdminLoginResponse struct {
	Token string
	Admin db.Management
}
