/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package request

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
