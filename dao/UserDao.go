/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package dao

import (
	"context"
	"errors"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserDaoInterface interface {
	GetGenderTotals() (int, int)
	TotalUsers() int
	UsersLocation() []string
	AddUser(user db.User) (db.User, error)
	UpdateUser(id string, key string, value string) (db.User, error)
	GetUserByPhone(phone string) (db.User, error)
	GetUserById(userId string) (db.User, error)
	DoesUserExist(phone string) bool
	VerifyUser(userid string) error
}

type NewUserDaoInterface struct {
	Client *mongo.Client
}

func (s NewUserDaoInterface) GetGenderTotals() (int, int) {
	var users []db.User
	var male = 0
	var female = 0
	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	cursor, err := MbungeDb.Find(context.Background(), bson.M{})
	if err != nil {
		return 0, 0
	}
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return 0, 0
	}
	for i := 0; i < len(users); i++ {
		if users[i].Gender == 0 {
			male = +1
		} else {
			female = +1
		}
	}
	return male, female
}
func (s NewUserDaoInterface) TotalUsers() int {
	var users []db.User
	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	cursor, err := MbungeDb.Find(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return 0
	}
	return len(users)
}

func findUserById(id primitive.ObjectID, client *mongo.Client) db.User {
	var user db.User
	MbungeDb := client.Database("mbunge").Collection("user")
	err := MbungeDb.FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&user)
	if err != nil {
		return db.User{}
	}
	return user
}
func (s NewUserDaoInterface) AddUser(user db.User) (db.User, error) {
	user.Verified = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectID()

	hashedPassword, err := utils.GenerateHash(user.Password)
	if err != nil {
		return db.User{}, err
	}
	user.Password = hashedPassword

	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	exist := s.DoesUserExist(user.PhoneNumber)
	if exist {
		return db.User{}, errors.New("user exist")
	} else {
		res, err := MbungeDb.InsertOne(context.Background(), user)
		if err != nil {
			return db.User{}, err
		}

		return findUserById(res.InsertedID.(primitive.ObjectID), s.Client), nil
	}
}
func (s NewUserDaoInterface) UpdateUser(id string, key string, value string) (db.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}
	//update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())}}}}
	MbungeDb := *s.Client.Database("mbunge").Collection("user")

	_, err := MbungeDb.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return db.User{}, err
	}

	return findUserById(objID, s.Client), nil
}
func (s NewUserDaoInterface) GetUserByPhone(phone string) (db.User, error) {
	var user db.User
	MbungeDb := *s.Client.Database("mbunge").Collection("user")

	err := MbungeDb.FindOne(context.Background(), bson.M{
		"phone_number": phone,
	}).Decode(&user)

	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (s NewUserDaoInterface) GetUserById(userId string) (db.User, error) {
	var user db.User

	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	objectID, _ := primitive.ObjectIDFromHex(userId)
	err := MbungeDb.FindOne(context.Background(), bson.M{
		"_id": objectID,
	}).Decode(&user)
	if err != nil {
		return db.User{}, nil
	}
	return user, nil
}

func (s NewUserDaoInterface) DoesUserExist(phone string) bool {
	var result bson.M
	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	err := MbungeDb.FindOne(context.Background(), bson.M{
		"phone_number": phone,
	}).Decode(&result)

	if err != nil {
		return false
	}
	if len(result) != 0 {
		return true
	} else {
		return false
	}
}
func (s NewUserDaoInterface) VerifyUser(userid string) error {

	objID, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.D{{"verified", true}}}}
	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	_, err := MbungeDb.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}
func (s NewUserDaoInterface) UsersLocation() []string {
	var locations []string
	var users []db.User
	MbungeDb := *s.Client.Database("mbunge").Collection("user")
	cursor, err := MbungeDb.Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil
	}
	for _, user := range users {
		if user.County != "" {
			locations = append(locations, user.County)
		}
	}
	return locations
}
