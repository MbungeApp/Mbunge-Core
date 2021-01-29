/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Management struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name" form:"name"`
	NationalID   string             `json:"national_id" bson:"national_id" form:"national_id"`
	EmailAddress string             `json:"email_address" bson:"email_address" form:"email_address"`
	Password     string             `json:"password" bson:"password" form:"password"`
	Role         int                `json:"role" bson:"role"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at" form:"updated"`
}