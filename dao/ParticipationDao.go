/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ParticipationDaoInterface interface {
	GetAllParticipation() []db.Participation
	GetParticipationByID() db.Participation
	ParticipationChanges() (*mongo.ChangeStream, error)
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

func (p NewParticipationDaoInterface) ParticipationChanges() (*mongo.ChangeStream, error) {
	matchStage1 := bson.D{{"$match", bson.D{{"operationType", "insert"}}}}
	matchStage2 := bson.D{{"$match", bson.D{{"operationType", "update"}}}}
	matchStage3 := bson.D{{"$match", bson.D{{"operationType", "delete"}}}}

	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	changeStream, err := participationCollection(p.Client).Watch(context.TODO(), mongo.Pipeline{
		matchStage1, matchStage2, matchStage3,
	}, opts)
	if err != nil {
		return nil, err
	}
	return changeStream, nil
}

//func (p NewParticipationDaoInterface) GetParticipationByID() db.Participation  {
//
//}
