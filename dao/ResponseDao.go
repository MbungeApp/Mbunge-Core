package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResponseDao interface {
	AddResponse(response db.Response) error
	GetAllResponseByParti(participationID string) []db.Participation
	DeleteResponse(responseId string) error
	MostResponse()
}

type NewResponseDao struct {
	Client *mongo.Client
}

func responseCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("response")
}
func (r NewResponseDao) AddResponse(response db.Response) error {

	response.ID = primitive.NewObjectID()
	_, err := responseCollection(r.Client).InsertOne(context.Background(), response)
	if err != nil {
		return err
	}
	return nil
}

func (r NewResponseDao) GetAllResponseByParti(participationID string) []db.Participation {
	var participation []db.Participation

	cursor, err := responseCollection(r.Client).Find(context.Background(), bson.D{
		{"participation_id", participationID},
	})
	if err != nil {
		return nil
	}
	err = cursor.All(context.Background(), &participation)
	if err != nil {
		return nil
	}
	return participation
}

func (r NewResponseDao) DeleteResponse(responseId string) error {
	objectID, _ := primitive.ObjectIDFromHex(responseId)
	_, err := responseCollection(r.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}
