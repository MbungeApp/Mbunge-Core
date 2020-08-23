package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipationService interface {
	GetAllParticipation() []db.Participation
	AddResponse(res request.ResponseRequest) error
	GetAllResponseByParti(participationId string) []db.Participation
	DeleteResponse(responseId string) error
	ChangeStreams() (*mongo.ChangeStream, error)
}