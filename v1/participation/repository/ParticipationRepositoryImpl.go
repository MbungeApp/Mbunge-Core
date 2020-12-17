/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type participationRepositoryImpl struct {
	participationDao dao.NewParticipationDaoInterface
	responseDao      dao.NewResponseDao
	userDao          dao.NewUserDaoInterface
}

func NewParticipationRepositoryImpl(client *mongo.Client) ParticipationRepository {
	participationDao := dao.NewParticipationDaoInterface{Client: client}
	responseDao := dao.NewResponseDao{Client: client}
	userDao := dao.NewUserDaoInterface{Client: client}

	return &participationRepositoryImpl{
		participationDao: participationDao,
		responseDao:      responseDao,
		userDao:          userDao,
	}
}

// Implemented methods
func (p participationRepositoryImpl) GetAllParticipation() []db.Participation {
	return p.participationDao.GetAllParticipation()
}

func (p participationRepositoryImpl) GetResponsesByParticipationID(participationID string) []db.Response {
	return p.responseDao.GetAllResponseByParti(participationID)
}

func (p participationRepositoryImpl) AddResponses(resp db.Response) (response.AddResponseResponse, error) {
	res, err := p.responseDao.AddResponse(resp)
	if err != nil {
		return response.AddResponseResponse{}, err
	}
	user, err := p.userDao.GetUserById(res.UserId)
	if err != nil {
		return response.AddResponseResponse{}, err
	}
	result := response.AddResponseResponse{
		ID:              res.ID,
		UserId:          res.UserId,
		ParticipationId: res.ParticipationId,
		Body:            res.Body,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
		User:            user,
	}
	return result, nil
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
