package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    string             `json:"user_id" bson:"user_id"`
	Type      string             `json:"type" bson:"type"`
	FCMToken  string             `json:"fcm_token" bson:"fcm_token"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
