/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ParticipationDaoInterface interface {
	TotalParticipations() int
	GetAllParticipation() []db.Participation
	GetParticipationByID(participationID string) (db.Participation, error)
	ParticipationChanges() (*mongo.ChangeStream, error)
	CreateParticipation(participation db.Participation) error
	UpdateParticipation(id string, key string, value string) error
	DeleteParticipation(participationID string) error
}

type NewParticipationDaoInterface struct {
	Client *mongo.Client
}

func participationCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("participation")
}

func (p NewParticipationDaoInterface) TotalParticipations() int {
	var participation []db.Participation
	cursor, err := participationCollection(p.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	err = cursor.All(context.Background(), &participation)
	if err != nil {
		return 0
	}
	return len(participation)
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

func (p NewParticipationDaoInterface) GetParticipationByID(participationID string) (db.Participation, error) {
	objectID, _ := primitive.ObjectIDFromHex(participationID)
	var participation db.Participation

	err := participationCollection(p.Client).FindOne(context.Background(), bson.D{{"_id", objectID}}).Decode(&participation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return db.Participation{}, err
		}
		return db.Participation{}, err
	}
	return participation, nil

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
func (p NewParticipationDaoInterface) CreateParticipation(participation db.Participation) error {
	participation.CreatedAt = time.Now()
	participation.UpdatedAt = time.Now()
	participation.ID = primitive.NewObjectID()
	_, err := participationCollection(p.Client).InsertOne(context.Background(), participation)

	if err != nil {
		return err
	}
	return nil
}
func (p NewParticipationDaoInterface) UpdateParticipation(id string, key string, value string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}

	_, err := participationCollection(p.Client).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}
func (p NewParticipationDaoInterface) DeleteParticipation(participationID string) error {
	objectID, _ := primitive.ObjectIDFromHex(participationID)
	_, err := participationCollection(p.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}
