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
	"time"
)

type WebinarDaoInterface interface {
	TotalWebinars() int
	GetAllWebinars() []db.Webinar
	GetWebinarsByID(webinarID string) (db.Webinar, error)
	CreateWebinars(webinar db.Webinar) error
	UpdateWebinars(id string, key string, value string) error
	DeleteWebinars(webinarId string) error
}

type NewWebinarDaoInterface struct {
	Client *mongo.Client
}

func webinarCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("webinar")
}
func (n NewWebinarDaoInterface) TotalWebinars() int {
	var webinars []db.Webinar
	cursor, err := webinarCollection(n.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	err = cursor.All(context.Background(), &webinars)
	if err != nil {
		return 0
	}
	return len(webinars)
}
func (n NewWebinarDaoInterface) GetAllWebinars() []db.Webinar {
	var webinars []db.Webinar
	cursor, err := webinarCollection(n.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}
	err = cursor.All(context.Background(), &webinars)
	if err != nil {
		return nil
	}
	return webinars
}
func (n NewWebinarDaoInterface) GetWebinarsByID(webinarID string) (db.Webinar, error) {
	objectID, _ := primitive.ObjectIDFromHex(webinarID)
	var webinar db.Webinar

	err := webinarCollection(n.Client).FindOne(context.Background(), bson.D{{"_id", objectID}}).Decode(&webinar)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return db.Webinar{}, err
		}
		return db.Webinar{}, err
	}
	return webinar, nil
}
func (n NewWebinarDaoInterface) CreateWebinars(webinar db.Webinar) error {
	webinar.CreatedAt = time.Now()
	webinar.UpdatedAt = time.Now()
	webinar.ID = primitive.NewObjectID()
	_, err := webinarCollection(n.Client).InsertOne(context.Background(), webinar)

	if err != nil {
		return err
	}
	return nil
}
func (n NewWebinarDaoInterface) UpdateWebinars(id string, key string, value string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}

	_, err := webinarCollection(n.Client).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}
func (n NewWebinarDaoInterface) DeleteWebinars(webinarId string) error {
	objectID, _ := primitive.ObjectIDFromHex(webinarId)
	_, err := webinarCollection(n.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}
