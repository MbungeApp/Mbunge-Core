package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsDaoInterface interface {
	ReadNews() ([]db.EventNew, error)
}

type NewEventDaoInterface struct {
	Client *mongo.Client
}

func eventCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("event")
}
func (u NewEventDaoInterface) ReadNews() ([]db.EventNew, error) {
	var events []db.EventNew

	cursor, err := eventCollection(u.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
