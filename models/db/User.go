package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FirstName    string             `json:"first_name" bson:"first_name" form:"first_name"`
	LastName     string             `json:"last_name" bson:"last_name" form:"last_name"`
	EmailAddress string             `json:"email_address" bson:"email_address" form:"email_address"`
	Password     string             `json:"password" bson:"password" form:"password"`
	PhoneNumber  string             `json:"phone_number" bson:"phone_number" form:"phone_number"`
	DateBirth    time.Time          `json:"date_birth" bson:"date_birth" form:"date_birth"`
	ProfilePic   string             `json:"profile_pic" bson:"profile_pic" form:"profile_pic"`
	Gender       int                `json:"gender" bson:"gender" form:"gender"`
	Verified     bool               `json:"verified" bson:"verified"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
