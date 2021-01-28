/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/v1/webinar/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type participationServiceImpl struct {
	participationRepository repository.ParticipationRepository
}

func NewParticipationServiceImpl(participationRepo repository.ParticipationRepository) ParticipationService {
	return &participationServiceImpl{participationRepository: participationRepo}
}

// Implemented methods
func (p participationServiceImpl) GetAllParticipation() []db.Participation {
	return p.participationRepository.GetAllParticipation()
}

func (p participationServiceImpl) AddResponse(res request.ResponseRequest) (response.AddResponseResponse, error) {

	resx := db.Response{
		UserId:          res.UserId,
		ParticipationId: res.ParticipationId,
		Body:            res.Body,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	callback, err := p.participationRepository.AddResponses(resx)
	if err != nil {
		return response.AddResponseResponse{}, err
	}
	return callback, nil
}

func (p participationServiceImpl) GetAllResponseByParti(participationId string) []response.AddResponseResponse {
	return p.participationRepository.GetResponsesByParticipationID(participationId)
}

func (p participationServiceImpl) DeleteResponse(responseId string) error {
	return p.participationRepository.DeleteResponses(responseId)
}

func (p participationServiceImpl) ChangeStreams() (*mongo.ChangeStream, error) {
	stream, err := p.participationRepository.DBChanges()
	if err != nil {
		return nil, err
	}
	return stream, nil
}
