package service

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

type dashboardServiceImpl struct {
	participationDao dao.ParticipationDaoInterface
	responseDao      dao.ResponseDaoInterface
	usersDao         dao.UserDaoInterface
	newsEventsDao    dao.NewsDaoInterface
}

// NewDashboardServiceImpl ..
func NewDashboardServiceImpl(client *mongo.Client) DashboardServices {
	partiDao := dao.NewParticipationDaoInterface{
		Client: client,
	}
	resDao := dao.NewResponseDaoInterface{
		Client: client,
	}
	newDao := dao.NewEventDaoInterface{
		Client: client,
	}
	userDao := dao.NewUserDaoInterface{Client: client}

	return &dashboardServiceImpl{
		participationDao: partiDao,
		responseDao:      resDao,
		usersDao:         userDao,
		newsEventsDao:    newDao,
	}
}
func (d dashboardServiceImpl) GetMetrics() response.Metrics {
	var geoCodeLocations []response.UserLocation
	card := response.Card{
		TotalUsers:         d.usersDao.TotalUsers(),
		TotalParticipation: d.participationDao.TotalParticipations(),
		TotalResponses:     d.responseDao.TotalResponses(),
		TotalEvents:        d.newsEventsDao.TotalNews(),
	}

	male, female := d.usersDao.GetGenderTotals()
	gender := response.GenderRation{
		Male:   male,
		Female: female,
	}

	userLocations := d.usersDao.UsersLocation()
	for _, element := range userLocations {
		//if contains(geoCodeLocations[], element) {
		//	fmt.Println("**************** todo ***************")
		//} else {
		lat, long := utils.LocationToGeoCode(element)
		location := response.UserLocation{
			Name:      element,
			Count:     0,
			Latitude:  lat,
			Longitude: long,
		}

		geoCodeLocations = append(geoCodeLocations, location)
		//	}
	}

	metrics := response.Metrics{
		Card:          card,
		GenderRation:  gender,
		MpOfTheWeek:   response.MpOfTheWeek{},
		UsersLocation: geoCodeLocations,
	}
	return metrics
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

// EVENTS

func (d dashboardServiceImpl) ViewAllEvents() ([]db.EventNew, error) {
	events, err := d.newsEventsDao.ReadNews()
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (d dashboardServiceImpl) ViewEventById(id string) db.EventNew {
	events := d.newsEventsDao.ReadOneNews(id)
	return events
}
func (d dashboardServiceImpl) AddEvent(event *request.EventRequest) error {
	eventDb := db.EventNew{
		Name:    event.Name,
		Body:    event.Body,
		Picture: event.Picture,
	}
	err := d.newsEventsDao.CreateNews(eventDb)
	if err != nil {
		return err
	}
	return nil
}
func (d dashboardServiceImpl) EditEvent(id string, event *request.EventRequest) error {
	originalEvent := d.newsEventsDao.ReadOneNews(id)

	if event.Name != originalEvent.Name {
		err := d.newsEventsDao.UpdateNews(id, "name", event.Name)
		if err != nil {
			return err
		}
	} else if event.Picture != originalEvent.Picture {
		err := d.newsEventsDao.UpdateNews(id, "picture", event.Picture)
		if err != nil {
			return err
		}
	} else if event.Body != originalEvent.Body {
		err := d.newsEventsDao.UpdateNews(id, "body", event.Body)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("******************** nothing *******************")
	}

	return nil
}
func (d dashboardServiceImpl) DeleteEvent(id string) error {
	err := d.newsEventsDao.DeleteNews(id)
	if err != nil {
		return err
	}
	return nil
}
