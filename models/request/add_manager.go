package request

type AddManager struct {
	Name         string `json:"name" bson:"name" form:"name"`
	NationalID   string `json:"national_id" bson:"national_id" form:"national_id"`
	EmailAddress string `json:"email_address" bson:"email_address" form:"email_address"`
	Role         int    `json:"role" bson:"role"`
}
