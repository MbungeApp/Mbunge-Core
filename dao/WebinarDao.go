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
	"strconv"
	"time"
)

type WebinarDaoInterface interface {
	TotalWebinars() int
	GetAllWebinars() ([]db.Webinar, error)
	GetWebinarsByID(webinarID string) (db.Webinar, error)
	CreateWebinars(webinar db.Webinar) error
	UpdateWebinars(id string, key string, value string) error
	DeleteWebinars(webinarId string) error
	WebinarsChanges() (*mongo.ChangeStream, error)
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
func (n NewWebinarDaoInterface) GetAllWebinars() ([]db.Webinar, error) {
	var webinars []db.Webinar
	cursor, err := webinarCollection(n.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &webinars)
	if err != nil {
		return nil, err
	}
	return webinars, nil
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
	webinar.Postponed = false
	webinar.ID = primitive.NewObjectID()
	_, err := webinarCollection(n.Client).InsertOne(context.Background(), webinar)

	if err != nil {
		return err
	}
	return nil
}
func (n NewWebinarDaoInterface) UpdateWebinars(id string, key string, value string) error {
	var update bson.D
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	if key == "schedule_at" {
		layout := "2006-01-02T15:04:05Z"
		parsedDOB, _ := time.Parse(layout, value)
		update = bson.D{{Key: "$set", Value: bson.M{key: parsedDOB, "updated_at": time.Now()}}}
	} else if key == "postponed" {
		result, _ := strconv.ParseBool(value)
		update = bson.D{{Key: "$set", Value: bson.M{key: result, "updated_at": time.Now()}}}
	} else if key == "duration" {
		i, _ := strconv.Atoi(value)
		update = bson.D{{Key: "$set", Value: bson.M{key: i, "updated_at": time.Now()}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}
	}

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

func (n NewWebinarDaoInterface) WebinarsChanges() (*mongo.ChangeStream, error) {

	changeStream, err := webinarCollection(n.Client).Watch(context.Background(), mongo.Pipeline{})
	if err != nil {
		return nil, err
	}
	return changeStream, nil
}
