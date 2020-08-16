package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipationDaoInterface interface {
	GetAllParticipation() []db.Participation
	GetParticipationByID() db.Participation
}

type NewParticipationDaoInterface struct {
	Client *mongo.Client
}

func participationCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("participation")
}

func (p NewParticipationDaoInterface) GetAllParticipation() []db.Participation {
	var participation []db.Participation
	cursor, err := participationCollection(p.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}
	err = cursor.All(context.Background(), &participation)
	if err != nil {
		return nil
	}
	return participation
}

//func (p NewParticipationDaoInterface) GetParticipationByID() db.Participation  {
//
//}
