/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package response

import "github.com/MbungeApp/mbunge-core/models/db"

type LoginResponse struct {
	Token string
	User  db.User
}
