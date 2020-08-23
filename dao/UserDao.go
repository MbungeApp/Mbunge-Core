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
	AddUser(user db.User) (db.User, error)
	UpdateUser(id string, key string, value string) (db.User, error)
	GetUserByPhone(phone string) (db.User, error)
	DoesUserExist(phone string) bool
	VerifyUser(userid string) error
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

	hashedPassword, err := utils.GenerateHash(user.Password)
	if err != nil {
		return db.User{}, err
	}
	user.Password = hashedPassword

	MbungeDb := *u.Client.Database("mbunge").Collection("user")
	exist := u.DoesUserExist(user.PhoneNumber)
	if exist {
		return db.User{}, errors.New("user exist")
	} else {
		res, err := MbungeDb.InsertOne(context.Background(), user)
		if err != nil {
			return db.User{}, err
		}

		return findUserById(res.InsertedID.(primitive.ObjectID), u.Client), nil
	}
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
func (u NewUserDaoInterface) GetUserByPhone(phone string) (db.User, error) {
	var user db.User
	MbungeDb := *u.Client.Database("mbunge").Collection("user")

	err := MbungeDb.FindOne(context.Background(), bson.M{
		"phone_number": phone,
	}).Decode(&user)

	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
func (u NewUserDaoInterface) DoesUserExist(phone string) bool {
	var result bson.M
	MbungeDb := *u.Client.Database("mbunge").Collection("user")
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
func (u NewUserDaoInterface) VerifyUser(userid string) error {

	objID, _ := primitive.ObjectIDFromHex(userid)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.D{{"verified", true}}}}
	MbungeDb := *u.Client.Database("mbunge").Collection("user")
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
