/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package dao

import (
	"context"
	"fmt"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type NewsDaoInterface interface {
	TotalNews() int
	ReadNews() ([]db.EventNew, error)
	ReadOneNews(newsID string) db.EventNew
	CreateNews(news db.EventNew) error
	UpdateNews(newsID string, key string, value string) error
	DeleteNews(newsID string) error
}

type NewEventDaoInterface struct {
	Client *mongo.Client
}

func eventCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("event")
}

//
func (n NewEventDaoInterface) TotalNews() int {
	var events []db.EventNew

	cursor, err := eventCollection(n.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	err = cursor.All(context.Background(), &events)
	if err != nil {
		return 0
	}
	return len(events)
}

//
func (n NewEventDaoInterface) ReadNews() ([]db.EventNew, error) {
	var events []db.EventNew

	cursor, err := eventCollection(n.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

//

func (n NewEventDaoInterface) ReadOneNews(newsID string) db.EventNew {
	objectID, _ := primitive.ObjectIDFromHex(newsID)
	var news db.EventNew

	err := eventCollection(n.Client).FindOne(context.Background(), bson.D{{"_id", objectID}}).Decode(&news)
	if err != nil {
		//log.Fatal(err)
		if err == mongo.ErrNoDocuments {
			return db.EventNew{}
		}
	}
	return news
}

//
func (n NewEventDaoInterface) CreateNews(news db.EventNew) error {
	news.UpdatedAt = time.Now()
	news.CreatedAt = time.Now()
	news.ID = primitive.NewObjectID()
	_, err := eventCollection(n.Client).InsertOne(context.Background(), news)
	if err != nil {
		return err
	}
	return nil
}

//
func (n NewEventDaoInterface) UpdateNews(newsID string, key string, value string) error {
	objID, _ := primitive.ObjectIDFromHex(newsID)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}

	result, err := eventCollection(n.Client).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return nil
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return nil
}

//
func (n NewEventDaoInterface) DeleteNews(newsID string) error {
	objectID, _ := primitive.ObjectIDFromHex(newsID)
	_, err := eventCollection(n.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil
}
