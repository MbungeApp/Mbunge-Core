/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package dao

import (
	"context"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MPDaoInterface interface {
	GetAllMps() ([]db.MP, error)
	GetMpOfTheWeek() response.MpOftheWeek
}

type NewMPDaoInterface struct {
	Client *mongo.Client
}

func mpCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("mbunge").Collection("mp")
}
func (m NewMPDaoInterface) GetAllMps() ([]db.MP, error) {
	var mps []db.MP

	cursor, err := mpCollection(m.Client).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &mps)
	if err != nil {
		return nil, err
	}
	return mps, nil
}
func (m NewMPDaoInterface) GetMpOfTheWeek() response.MpOftheWeek {
	var mpLive []db.MpLive
	var mp db.MP

	//
	dbTable := m.Client.Database("mbunge").Collection("mp-live")
	cursor, err := dbTable.Find(context.Background(), bson.M{})
	if err != nil {
		return response.MpOftheWeek{}
	}
	err = cursor.All(context.Background(), &mpLive)
	if err != nil {
		return response.MpOftheWeek{}
	}

	//
	err = responseCollection(m.Client).FindOne(context.Background(),
		bson.D{{"mp_id", mpLive[0].MpId}}).Decode(&mp)
	if err != nil {
		return response.MpOftheWeek{}
	}
	res := response.MpOftheWeek{
		Mp:      mp,
		Details: mpLive[0],
	}
	return res
}
