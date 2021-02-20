package request

import "time"

type EditUser struct {
	FirstName    string    `json:"first_name" bson:"first_name" form:"first_name"`
	LastName     string    `json:"last_name" bson:"last_name" form:"last_name"`
	EmailAddress string    `json:"email_address" bson:"email_address" form:"email_address"`
	PhoneNumber  string    `json:"phone_number" bson:"phone_number" form:"phone_number"`
	DateBirth    time.Time `json:"date_birth" bson:"date_birth" form:"date_birth"`
	Gender       int       `json:"gender" bson:"gender" form:"gender"`
	County       string    `json:"county" bson:"county" form:"county"`
}
