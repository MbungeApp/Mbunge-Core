/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/v1/mp/repository"
)

type newMpServiceImpl struct {
	mpRepo repository.MpRepository
}

func NewMpService(mpRepository repository.MpRepository) MpService {
	return &newMpServiceImpl{mpRepo: mpRepository}
}

func (n newMpServiceImpl) AllMps() ([]db.MP, error) {
	mps, err := n.mpRepo.GetAllMps()
	if err != nil {
		return nil, err
	}
	return mps, nil
}

func (n newMpServiceImpl) MpOftheWeek() response.MpOftheWeek {
	res := n.mpRepo.GetMpOftheWeek()
	return res
}
