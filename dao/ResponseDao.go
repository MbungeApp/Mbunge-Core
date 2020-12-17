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
)

type ResponseDao interface {
	AddResponse(response db.Response) error
	GetAllResponseByParti(participationID string) []db.Response
	GetResponseById(responseId string) (db.Response, error)
	DeleteResponse(responseId string) error
	MostResponse()
}

type NewResponseDao struct {
	Client *mongo.Client
}

func responseCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("response")
}
func (r NewResponseDao) AddResponse(response db.Response) (db.Response, error) {

	response.ID = primitive.NewObjectID()
	res, err := responseCollection(r.Client).InsertOne(context.Background(), response)
	if err != nil {
		return db.Response{}, err
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)
	response, err = r.GetResponseById(oid.Hex())
	if err != nil {
		return db.Response{}, nil
	}
	return response, nil
}

func (r NewResponseDao) GetAllResponseByParti(participationID string) []db.Response {
	var responses []db.Response

	cursor, err := responseCollection(r.Client).Find(context.Background(), bson.D{
		{"participation_id", participationID},
	})
	if err != nil {
		return []db.Response{}
	}
	err = cursor.All(context.Background(), &responses)
	if err != nil {
		return []db.Response{}
	}
	return responses
}

func (r NewResponseDao) GetResponseById(responseId string) (db.Response, error) {
	var response db.Response

	objectID, _ := primitive.ObjectIDFromHex(responseId)
	err := responseCollection(r.Client).FindOne(context.Background(), bson.M{
		"_id": objectID,
	}).Decode(&response)
	if err != nil {
		return db.Response{}, nil
	}
	return response, nil
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
