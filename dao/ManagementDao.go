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
	"strconv"
	"time"
)

type ManagementDaoInterface interface {

	// Users
	ReadManagers() []db.Management
	InsertManagers(user db.Management) error
	FindAdminById(id string) db.Management
	FindManagerByEmail(email string) (db.Management, error)
	//UpdateRole(userId string, roleId string) error
	UpdateManager(userId string, key string, value string) error
	DeleteManager(userId string) error
}

type NewManagementDaoInterface struct {
	Client *mongo.Client
}

func managersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("management")
}

// Managers
func (m NewManagementDaoInterface) ReadManagers() []db.Management {
	var managers []db.Management

	cursor, err := managersCollection(m.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}
	err = cursor.All(context.Background(), &managers)
	if err != nil {
		return nil
	}
	return managers
}
func (m NewManagementDaoInterface) InsertManagers(user db.Management) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	//user.Password = "admin1234"
	_, err := managersCollection(m.Client).InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (m NewManagementDaoInterface) UpdateManager(userId string, key string, value string) error {
	objID, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", objID}}
	var update interface{}
	if key == "role" {
		i, _ := strconv.Atoi(value)
		update = bson.D{{Key: "$set", Value: bson.M{key: i, "updated_at": time.Now()}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}
	}
	result, err := managersCollection(m.Client).UpdateOne(
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

func (m NewManagementDaoInterface) DeleteManager(userId string) error {
	objectID, _ := primitive.ObjectIDFromHex(userId)
	_, err := managersCollection(m.Client).DeleteOne(context.Background(), bson.M{
		"_id": objectID,
	})
	if err != nil {
		return err
	}
	return nil

}
func (m NewManagementDaoInterface) FindManagerByEmail(email string) (db.Management, error) {
	var manager db.Management

	err := managersCollection(m.Client).FindOne(context.Background(), bson.M{
		"email_address": email,
	}).Decode(&manager)

	if err != nil {
		return db.Management{}, err
	}
	return manager, nil
}
func (m NewManagementDaoInterface) FindAdminById(id string) db.Management {
	objectID, _ := primitive.ObjectIDFromHex(id)
	var manager db.Management

	err := managersCollection(m.Client).FindOne(context.Background(), bson.D{{"_id", objectID}}).Decode(&manager)
	if err != nil {
		//log.Fatal(err)
		if err == mongo.ErrNoDocuments {
			return db.Management{}
		}
	}
	return manager
}
