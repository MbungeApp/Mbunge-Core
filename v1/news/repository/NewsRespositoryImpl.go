package repository

import (
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type newsRepository struct {
	eventDao dao.NewEventDaoInterface
}

func NewEventRepository(client *mongo.Client) NewsRepository {
	eventDao := dao.NewEventDaoInterface{Client: client}
	return &newsRepository{eventDao: eventDao}
}

// Implemented methods
func (n newsRepository) GetAllNews() ([]db.EventNew, error) {
	newsEvents, err := n.eventDao.ReadNews()
	if err != nil {
		return nil, err
	}
	return newsEvents, nil
}
