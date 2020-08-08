package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DeviceDaoInterface interface {
	AddDevice(device db.Device) (db.Device, error)
	UpdateDevice(id string, key string, value string) (db.Device, error)
	GetDevice(userId string) (db.Device, error)
}
type NewDeviceDaoInterface struct {
	Client *mongo.Client
}

// Returns Device table
func deviceCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("device")
}
func findDeviceById(id primitive.ObjectID, client *mongo.Client) db.Device {
	var device db.Device
	err := deviceCollection(client).FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&device)
	if err != nil {
		return db.Device{}
	}
	return device
}

func (d NewDeviceDaoInterface) AddDevice(device db.Device) (db.Device, error) {
	device.ID = primitive.NewObjectID()
	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()
	res, err := deviceCollection(d.Client).InsertOne(context.Background(), device)
	if err != nil {
		return db.Device{}, err
	}
	return findDeviceById(res.InsertedID.(primitive.ObjectID), d.Client), nil
}
func (d NewDeviceDaoInterface) UpdateDevice(id string, key string, value string) (db.Device, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{Key: "$set", Value: bson.M{key: value, "updated_at": time.Now()}}}
	_, err := deviceCollection(d.Client).UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return db.Device{}, nil
	}
	return findDeviceById(objID, d.Client), nil
}
func (d NewDeviceDaoInterface) GetDevice(userId string) (db.Device, error) {
	var device db.Device

	err := deviceCollection(d.Client).FindOne(context.Background(), bson.M{
		"user_id": userId,
	}).Decode(&device)

	if err != nil {
		return db.Device{}, err
	}
	return device, nil
}
