package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
)

type WebinarService interface {
	GetAllWebinars() []db.Webinar
	AddResponse(res request.ResponseRequest) (response.AddResponseResponse, error)
	GetAllResponseByParti(participationId string) []response.AddResponseResponse
	DeleteResponse(responseId string) error
}
