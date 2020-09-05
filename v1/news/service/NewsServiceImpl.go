/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package service

import (
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/v1/news/repository"
)

type eventServiceImpl struct {
	newsRepo repository.NewsRepository
}

func NewEventService(eventRepository repository.NewsRepository) NewsService {
	return &eventServiceImpl{newsRepo: eventRepository}
}

// Implemented methods
func (e eventServiceImpl) AllNews() ([]db.EventNew, error) {
	events, err := e.newsRepo.GetAllNews()
	if err != nil {
		return nil, err
	}
	return events, nil
}
