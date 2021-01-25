package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Webinar struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Agenda      string             `json:"agenda" bson:"agenda"`
	HostedBy    string             `json:"hosted_by" bson:"hosted_by"`
	Description string             `json:"description" bson:"description"`
	Duration    int                `json:"duration" bson:"duration"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at" form:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at" form:"updated"`
	ScheduleAt  time.Time          `json:"schedule_at" bson:"schedule_at"`
	Postponed   bool               `json:"postponed" bson:"postponed"`
}
