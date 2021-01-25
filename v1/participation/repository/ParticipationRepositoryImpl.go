/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package repository

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type participationRepositoryImpl struct {
	participationDao dao.NewParticipationDaoInterface
	responseDao      dao.NewResponseDaoInterface
	userDao          dao.NewUserDaoInterface
}

func NewParticipationRepositoryImpl(client *mongo.Client) ParticipationRepository {
	participationDao := dao.NewParticipationDaoInterface{Client: client}
	responseDao := dao.NewResponseDaoInterface{Client: client}
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

func (p participationRepositoryImpl) GetResponsesByParticipationID(participationID string) []response.AddResponseResponse {
	fmt.Println("Shit 1")
	var results []response.AddResponseResponse
	responses := p.responseDao.GetAllResponseByParti(participationID)
	fmt.Println("Shit 2")
	for index, element := range responses {
		fmt.Println(index)
		user, _ := p.userDao.GetUserById(element.UserId)
		appendWhat := response.AddResponseResponse{
			ID:              element.ID,
			UserId:          element.UserId,
			ParticipationId: element.ParticipationId,
			Body:            element.Body,
			CreatedAt:       element.CreatedAt,
			UpdatedAt:       element.UpdatedAt,
			User:            user,
		}
		results = append(results, appendWhat)
	}
	return results
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
