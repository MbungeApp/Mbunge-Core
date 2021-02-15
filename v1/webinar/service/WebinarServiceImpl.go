package service

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type webinarServiceImpl struct {
	responseDao dao.ResponseDaoInterface
	webinarDao  dao.WebinarDaoInterface
	usersDao    dao.UserDaoInterface
}

// NewWebinarServiceImpl ..
func NewWebinarServiceImpl(client *mongo.Client) WebinarService {
	userDao := dao.NewUserDaoInterface{Client: client}
	resDao := dao.NewResponseDaoInterface{Client: client}
	webinarDao := dao.NewWebinarDaoInterface{Client: client}
	return &webinarServiceImpl{
		responseDao: resDao,
		webinarDao:  webinarDao,
		usersDao:    userDao,
	}
}

func (w webinarServiceImpl) GetAllWebinars() []db.Webinar {
	webinars, err := w.webinarDao.GetAllWebinars()
	if err != nil {
		return nil
	}
	return webinars
}

func (w webinarServiceImpl) AddResponse(res request.ResponseRequest) (response.AddResponseResponse, error) {
	resx := db.Response{
		UserId:          res.UserId,
		ParticipationId: res.ParticipationId,
		Body:            res.Body,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	resp, err := w.responseDao.AddResponse(resx)
	if err != nil {
		return response.AddResponseResponse{}, err
	}
	user, err := w.usersDao.GetUserById(res.UserId)
	if err != nil {
		return response.AddResponseResponse{}, err
	}
	result := response.AddResponseResponse{
		ID:              resp.ID,
		UserId:          resp.UserId,
		ParticipationId: resp.ParticipationId,
		Body:            resp.Body,
		CreatedAt:       resp.CreatedAt,
		UpdatedAt:       resp.UpdatedAt,
		User:            user,
	}
	return result, nil
}

func (w webinarServiceImpl) GetAllResponseByParti(participationId string) []response.AddResponseResponse {
	fmt.Println("Shit 1")
	var results []response.AddResponseResponse
	responses := w.responseDao.GetAllResponseByParti(participationId)
	fmt.Println("Shit 2")
	for index, element := range responses {
		fmt.Println(index)
		user, _ := w.usersDao.GetUserById(element.UserId)
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

func (w webinarServiceImpl) DeleteResponse(responseId string) error {
	err := w.responseDao.DeleteResponse(responseId)
	if err != nil {
		return err
	}
	return nil
}

func (w webinarServiceImpl) ChangeStreams() (*mongo.ChangeStream, error) {
	return w.responseDao.QuestionChanges()
}
