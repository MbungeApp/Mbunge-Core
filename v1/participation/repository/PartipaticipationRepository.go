package repository

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipationRepository interface {
	GetAllParticipation() []db.Participation
	GetResponsesByParticipationID(participationID string) []db.Participation
	AddResponses(response db.Response) error
	DeleteResponses(responseId string) error
	DBChanges() (*mongo.ChangeStream, error)
}
