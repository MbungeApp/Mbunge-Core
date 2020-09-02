package repository

import (
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type mpRepository struct {
	mpDao dao.NewMPDaoInterface
}

func NewMpRepository(client *mongo.Client) MpRepository {
	return &mpRepository{mpDao: dao.NewMPDaoInterface{Client: client}}
}

//
func (m mpRepository) GetAllMps() ([]db.MP, error) {
	mps, err := m.mpDao.GetAllMps()
	if err != nil {
		return nil, err
	}
	return mps, nil
}

func (m mpRepository) GetMpOftheWeek() response.MpOftheWeek {
	res := m.mpDao.GetMpOfTheWeek()
	return res
}
