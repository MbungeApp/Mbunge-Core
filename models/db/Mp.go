package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MP struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Image         string             `json:"image" bson:"image"`
	Constituency  string             `json:"constituency" bson:"constituency"`
	County        string             `json:"county" bson:"county"`
	MartialStatus string             `json:"martial_status" bson:"martial_status"`
	DateBirth     time.Time          `json:"date_birth" bson:"date_birth"`
	Bio           string             `json:"bio" bson:"bio"`
	Images        []string           `json:"images" bson:"images"`
}