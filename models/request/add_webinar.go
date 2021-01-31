package request

import "time"

type AddWebinar struct {
	Agenda      string    `json:"agenda" bson:"agenda"`
	HostedBy    string    `json:"hosted_by" bson:"hosted_by"`
	Description string    `json:"description" bson:"description"`
	Duration    int       `json:"duration" bson:"duration"`
	ScheduleAt  time.Time `json:"schedule_at" bson:"schedule_at"`
}
