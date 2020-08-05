package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserDaoInterface interface {
	AddUser(user db.User) (db.User, error)
	UpdateUser(id string, key string, value string) (db.User, error)
	GetUserByEmailAddress(emailAddress string) (db.User, error)
}
type NewUserDaoInterface struct {
	Client *mongo.Client
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
func (u NewUserDaoInterface) AddUser(user db.User) (db.User, error) {
	user.Verified = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectID()
	MbungeDb := *u.Client.Database("mbunge").Collection("user")

	res, err := MbungeDb.InsertOne(context.Background(), user)
	if err != nil {
		return db.User{}, err
	}

	return findUserById(res.InsertedID.(primitive.ObjectID), u.Client), nil
}
func (u NewUserDaoInterface) UpdateUser(id string, key string, value string) (db.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}
	//update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": primitive.Timestamp{T: uint32(time.Now().Unix())}}}}
	MbungeDb := *u.Client.Database("mbunge").Collection("user")

	_, err := MbungeDb.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return db.User{}, err
	}

	return findUserById(objID, u.Client), nil
}
func (u NewUserDaoInterface) GetUserByEmailAddress(emailAddress string) (db.User, error) {
	var user db.User
	MbungeDb := *u.Client.Database("mbunge").Collection("user")

	err := MbungeDb.FindOne(context.Background(), bson.M{
		"email_address": emailAddress,
	}).Decode(&user)

	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
