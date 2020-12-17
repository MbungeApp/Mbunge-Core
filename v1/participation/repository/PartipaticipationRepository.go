/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipationRepository interface {
	GetAllParticipation() []db.Participation
	GetResponsesByParticipationID(participationID string) []db.Response
	AddResponses(response db.Response) (response.AddResponseResponse, error)
	DeleteResponses(responseId string) error
	DBChanges() (*mongo.ChangeStream, error)
}
