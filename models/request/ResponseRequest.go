/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package request

type ResponseRequest struct {
	UserId          string `json:"user_id"`
	ParticipationId string `json:"participation_id"`
	Body            string `json:"body"`
}
