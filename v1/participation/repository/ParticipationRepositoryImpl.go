/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type participationRepositoryImpl struct {
	participationDao dao.NewParticipationDaoInterface
	responseDao      dao.NewResponseDao
}

func NewParticipationRepositoryImpl(client *mongo.Client) ParticipationRepository {
	participationDao := dao.NewParticipationDaoInterface{Client: client}
	responseDao := dao.NewResponseDao{Client: client}

	return &participationRepositoryImpl{
		participationDao: participationDao,
		responseDao:      responseDao,
	}
}

// Implemented methods
func (p participationRepositoryImpl) GetAllParticipation() []db.Participation {
	return p.participationDao.GetAllParticipation()
}

func (p participationRepositoryImpl) GetResponsesByParticipationID(participationID string) []db.Participation {
	return p.responseDao.GetAllResponseByParti(participationID)
}

func (p participationRepositoryImpl) AddResponses(response db.Response) error {
	err := p.responseDao.AddResponse(response)
	if err != nil {
		return err
	}
	return nil
}

func (p participationRepositoryImpl) DeleteResponses(responseId string) error {
	err := p.responseDao.DeleteResponse(responseId)
	if err != nil {
		return err
	}
	return nil
}

func (p participationRepositoryImpl) DBChanges() (*mongo.ChangeStream, error) {
	stream, err := p.participationDao.ParticipationChanges()
	if err != nil {
		return nil, err
	}
	return stream, nil
}
