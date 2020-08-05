package db

type OtpCode struct {
	UserId string `json:"user_id" bson:"user_id"`
	Code   int    `json:"code" bson:"code"`
}
