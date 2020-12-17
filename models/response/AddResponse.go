package response

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AddResponseResponse struct {
	ID              primitive.ObjectID `json:"id" `
	UserId          string             `json:"user_id"`
	ParticipationId string             `json:"participation_id" `
	Body            string             `json:"body" bson:"body"`
	CreatedAt       time.Time          `json:"created_at" `
	UpdatedAt       time.Time          `json:"updated_at"`
	User            db.User            `json:"user"`
}
