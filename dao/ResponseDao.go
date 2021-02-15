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

type ResponseDaoInterface interface {
	TotalResponses() int
	AddResponse(response db.Response) (db.Response, error)
	GetAllResponseByParti(participationID string) []db.Response
	GetResponseById(responseId string) (db.Response, error)
	DeleteResponse(responseId string) error
	QuestionChanges() (*mongo.ChangeStream, error)
	//MostResponse()
}

type NewResponseDaoInterface struct {
	Client *mongo.Client
}

func responseCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("response")
}
func (r NewResponseDaoInterface) TotalResponses() int {
	var responses []db.Response
	cursor, err := responseCollection(r.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	err = cursor.All(context.Background(), &responses)
	if err != nil {
		return 0
	}

	return len(responses)
}
func (r NewResponseDaoInterface) AddResponse(response db.Response) (db.Response, error) {

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

func (r NewResponseDaoInterface) GetAllResponseByParti(participationID string) []db.Response {
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

func (r NewResponseDaoInterface) GetResponseById(responseId string) (db.Response, error) {
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

func (r NewResponseDaoInterface) DeleteResponse(responseId string) error {
	objectID, _ := primitive.ObjectIDFromHex(responseId)
	_, err := responseCollection(r.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r NewResponseDaoInterface) QuestionChanges() (*mongo.ChangeStream, error) {

	changeStream, err := responseCollection(r.Client).Watch(context.Background(), mongo.Pipeline{})
	if err != nil {
		return nil, err
	}
	return changeStream, nil
}
