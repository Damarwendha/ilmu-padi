package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseFavouriteControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *mocking.CourseFavouriteUsecaseMock
	amm *mocking.AuthMiddlewareMock
}

func (s *CourseFavouriteControllerTestSuite) SetupTest() {
	s.amm = new(mocking.AuthMiddlewareMock)
	s.aum = new(mocking.CourseFavouriteUsecaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")

	rg.Use(s.amm.RequireToken("user", "admin"))
	s.rg = rg
}

func TestCoursesCtrlTestSuite(t *testing.T) {
	suite.Run(t, new(CourseFavouriteControllerTestSuite))
}

func (s *CourseFavouriteControllerTestSuite) TestAddOrRemoveCourseFavourite_Success() {
	payload := entity.UserCoursesFavourite{
		UserID:   "1",
		CourseID: "1",
	}

	s.amm.On("RequireToken", "user", "admin").Return(nil)
	s.aum.On("AddOrRemoveToFavourite", payload).Return("Add/Remove Course Favourite Executed", nil)

	coursesCtrl := NewUserCoursesFavouriteController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	jsonPayload, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", "/api/v1/user/course/favourite", bytes.NewBuffer(jsonPayload))
	assert.NoError(s.T(), err)
	request.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = request

	coursesCtrl.AddOrRemoveCourseFavourite(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseFavouriteControllerTestSuite) TestAddOrRemoveCourseFavourite_Failed() {
	payload := entity.UserCoursesFavourite{
		UserID:   "1",
		CourseID: "1",
	}

	s.amm.On("RequireToken", "user", "admin").Return(nil)
	s.aum.On("AddOrRemoveToFavourite", payload).Return("", fmt.Errorf("Add/Remove Course Favourite Failed"))

	coursesCtrl := NewUserCoursesFavouriteController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	jsonPayload, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", "/api/v1/user/course/favourite", bytes.NewBuffer(jsonPayload))
	assert.NoError(s.T(), err)
	request.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = request

	coursesCtrl.AddOrRemoveCourseFavourite(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *CourseFavouriteControllerTestSuite) TestGetAllCourseFavourite_Success() {
	s.amm.On("RequireToken", "user", "admin").Return(nil)
	s.aum.On("GetAllFavouriteCourse", "1").Return([]entity.UserCoursesFavourite{
		{
			UserID:   "1",
			CourseID: "1",
		},
	}, nil)

	coursesCtrl := NewUserCoursesFavouriteController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	request, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/user/course/favourite/%v", "1"), nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = request

	coursesCtrl.GetUserFavouriteList(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}
