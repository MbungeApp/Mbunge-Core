package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/v1/participation/repository"
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

func (p participationServiceImpl) AddResponse(res request.ResponseRequest) error {
	response := db.Response{
		UserId:          res.UserId,
		ParticipationId: res.ParticipationId,
		Body:            res.Body,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := p.participationRepository.AddResponses(response)
	if err != nil {
		return err
	}
	return nil
}

func (p participationServiceImpl) GetAllResponseByParti(participationId string) []db.Participation {
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
