package service

import (
	"fmt"
	"github.com/MbungeApp/mbunge-core/dao"
	"github.com/MbungeApp/mbunge-core/models/db"
	"github.com/MbungeApp/mbunge-core/models/request"
	"github.com/MbungeApp/mbunge-core/models/response"
	"github.com/MbungeApp/mbunge-core/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
	"time"
)

type dashboardServiceImpl struct {
	participationDao dao.ParticipationDaoInterface
	responseDao      dao.ResponseDaoInterface
	usersDao         dao.UserDaoInterface
	newsEventsDao    dao.NewsDaoInterface
	mpDao            dao.MPDaoInterface
	managerDao       dao.ManagementDaoInterface
	webinarDao       dao.WebinarDaoInterface
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
	mpDao := dao.NewMPDaoInterface{
		Client: client,
	}
	managerDao := dao.NewManagementDaoInterface{
		Client: client,
	}
	userDao := dao.NewUserDaoInterface{Client: client}
	webinarDao := dao.NewWebinarDaoInterface{Client: client}

	return &dashboardServiceImpl{
		participationDao: partiDao,
		responseDao:      resDao,
		usersDao:         userDao,
		newsEventsDao:    newDao,
		mpDao:            mpDao,
		managerDao:       managerDao,
		webinarDao:       webinarDao,
	}
}

// ****************************
// Metrics
// ****************************
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

// ****************************
// Webinar
// ****************************
func (d dashboardServiceImpl) ViewAllWebinars() ([]db.Webinar, error) {
	webinars, err := d.webinarDao.GetAllWebinars()
	if err != nil {
		return nil, err
	}
	return webinars, nil
}
func (d dashboardServiceImpl) AddWebinar(webinar *request.AddWebinar) error {
	webinarDb := db.Webinar{
		Agenda:      webinar.Agenda,
		HostedBy:    webinar.HostedBy,
		Description: webinar.Description,
		Duration:    webinar.Duration,
		ScheduleAt:  webinar.ScheduleAt,
	}
	err := d.webinarDao.CreateWebinars(webinarDb)
	if err != nil {
		return err
	}
	go utils.SendNotification("notifications", map[string]map[string]string{
		"notification": {
			"body":  webinar.Description,
			"title": fmt.Sprintf("New webinar: %s", webinar.Agenda),
		},
	})
	return nil
}
func (d dashboardServiceImpl) DeleteWebinar(id string) error {
	err := d.webinarDao.DeleteWebinars(id)
	if err != nil {
		return err
	}
	return nil
}

// ****************************
// EVENTS
// ****************************
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

// ****************************
// MPs
// ****************************
func (d dashboardServiceImpl) ViewAllMps() ([]db.MP, error) {
	mps, err := d.mpDao.GetAllMps()
	if err != nil {
		return nil, err
	}
	return mps, nil
}

func (d dashboardServiceImpl) ViewMpById(id string) db.MP {
	mp := d.mpDao.ReadOneMp(id)
	return mp
}

func (d dashboardServiceImpl) AddMp(mp *request.MpRequest) error {
	layout := "2006-01-02"
	parsedDOB, err := time.Parse(layout, mp.DateOfBirth)

	mpDb := db.MP{
		Name:          mp.Name,
		Image:         mp.Picture,
		Constituency:  mp.Constituency,
		County:        mp.County,
		MartialStatus: mp.MartialStatus,
		DateBirth:     parsedDOB,
		Bio:           mp.Bio,
		Images:        nil,
	}

	err = d.mpDao.CreateMP(mpDb)
	if err != nil {
		return err
	}
	return nil
}

func (d dashboardServiceImpl) EditMp(id string, mp *request.MpRequest) error {
	layout := "2006-01-02"
	parsedDOB, _ := time.Parse(layout, mp.DateOfBirth)
	originalMp := d.mpDao.ReadOneMp(id)

	if mp.Name != originalMp.Name {
		err := d.mpDao.UpdateMPs(id, "name", mp.Name)
		if err != nil {
			return nil
		}
	} else if mp.Bio != originalMp.Bio {
		err := d.mpDao.UpdateMPs(id, "bio", mp.Bio)
		if err != nil {
			return nil
		}
	} else if parsedDOB != originalMp.DateBirth {
		err := d.mpDao.UpdateMPs(id, "date_birth", mp.DateOfBirth)
		if err != nil {
			return nil
		}
	} else if mp.MartialStatus != originalMp.MartialStatus {
		err := d.mpDao.UpdateMPs(id, "martial_status", mp.MartialStatus)
		if err != nil {
			return nil
		}
	} else if mp.County != originalMp.County {
		err := d.mpDao.UpdateMPs(id, "county", mp.County)
		if err != nil {
			return nil
		}
	} else if mp.Constituency != originalMp.Constituency {
		err := d.mpDao.UpdateMPs(id, "constituency", mp.Constituency)
		if err != nil {
			return nil
		}
	} else if mp.Picture != originalMp.Image {
		err := d.mpDao.UpdateMPs(id, "image", mp.Picture)
		if err != nil {
			return nil
		}
	} else {
		fmt.Println("******************** nothing *******************")
	}
	return nil
}

func (d dashboardServiceImpl) DeleteMp(id string) error {
	err := d.mpDao.DeleteMPs(id)
	if err != nil {
		return err
	}
	return nil
}

// ****************************
// ADMIN
// ****************************
func (d dashboardServiceImpl) FetchAllAdmins() []db.Management {
	admins := d.managerDao.ReadManagers()
	return admins
}
func (d dashboardServiceImpl) RegisterAdmin(regRequest request.AddManager) error {
	password := utils.RandStringBytes(5)

	hashedPassword, err := utils.GenerateHash(password)
	if err != nil {
		return err
	}
	manager := db.Management{
		ID:           primitive.NewObjectID(),
		Name:         regRequest.Name,
		NationalID:   regRequest.NationalID,
		EmailAddress: regRequest.EmailAddress,
		Password:     hashedPassword,
		Role:         regRequest.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = d.managerDao.InsertManagers(manager)
	if err != nil {
		return err
	}
	fmt.Println(password)
	utils.SendRandomEmail(regRequest.EmailAddress, password)
	return nil
}
func (d dashboardServiceImpl) LoginAdmin(request request.AdminLoginRequest) (response.AdminLoginResponse, error) {
	email := request.Email
	password := request.Password

	user, err := d.managerDao.FindManagerByEmail(email)
	if err != nil {
		return response.AdminLoginResponse{}, err
	}
	match, err := utils.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		fmt.Println(err)
		return response.AdminLoginResponse{}, err
	}
	if match {
		// new token
		token, err := utils.GenerateToken(user.EmailAddress)
		if err != nil {
			return response.AdminLoginResponse{}, err
		}
		res := response.AdminLoginResponse{
			Token: token,
			Admin: user,
		}
		return res, nil
	} else {
		return response.AdminLoginResponse{}, nil
	}
}
func (d dashboardServiceImpl) EditAdmin(id string, admin *request.AddManager) error {
	originalAdmin := d.managerDao.FindAdminById(id)

	fmt.Printf("Got: %s\n", originalAdmin.Name)
	if admin.Name != originalAdmin.Name {
		err := d.managerDao.UpdateManager(id, "name", admin.Name)
		if err != nil {
			return err
		}
	} else if admin.NationalID != originalAdmin.NationalID {
		err := d.managerDao.UpdateManager(id, "national_id", admin.NationalID)
		if err != nil {
			return err
		}
	} else if admin.Role != originalAdmin.Role {
		err := d.managerDao.UpdateManager(id, "role", string(rune(admin.Role)))
		if err != nil {
			return err
		}
	} else {
		fmt.Println("******************** nothing *******************")
	}
	return nil
}

func (d dashboardServiceImpl) UpdateAdminPassword(password *request.UpdatePassword) (db.Management, error) {
	user, err := d.managerDao.FindManagerByEmail(password.EmailAddress)
	if err != nil {
		return db.Management{}, err
	}
	hashedPassword, err := utils.GenerateHash(password.NewPassword)
	if err != nil {
		return db.Management{}, err
	}
	err = d.managerDao.UpdateManager(user.ID.String(), "password", hashedPassword)
	if err != nil {
		return db.Management{}, err
	}
	originalAdmin := d.managerDao.FindAdminById(user.ID.String())

	return originalAdmin, nil
}

func (d dashboardServiceImpl) FetchAdminById(id string) db.Management {
	admin := d.managerDao.FindAdminById(id)
	return admin
}
func (d dashboardServiceImpl) DeleteAdmin(id string) error {
	err := d.managerDao.DeleteManager(id)
	if err != nil {
		return err
	}
	return nil
}
