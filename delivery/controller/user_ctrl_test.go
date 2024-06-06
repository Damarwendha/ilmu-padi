package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *mocking.UserUsecaseMock
	amm *mocking.AuthMiddlewareMock
}

func (s *UserControllerTestSuite) SetupTest() {
	s.amm = new(mocking.AuthMiddlewareMock)
	s.aum = new(mocking.UserUsecaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")
	rg.Use(s.amm.RequireToken("user", "admin"))
	s.rg = rg
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (s *UserControllerTestSuite) TestGetAllProfile_Success() {
	s.aum.On("FindAll", 1, 10, "1").Return([]entity.User{}, model.Paging{}, nil)

	ctrl := NewUserController(s.aum, s.rg, s.amm)
	ctrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/profile", nil)
	assert.NoError(s.T(), err)
	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	ctrl.GetAllProfile(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *UserControllerTestSuite) TestGetProfileByID() {
	// Mocking userUsecase.GetProfileByID
	id := uint(1)
	user := "user"
	userExpected := entity.User{}
	s.aum.On("GetProfileByID", id, user).Return(userExpected, nil)

	// Create a new UserController
	ctrl := NewUserController(s.aum, s.rg, s.amm)
	ctrl.Route()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/profile/1", nil)
	assert.NoError(s.T(), err)
	record := httptest.NewRecorder()

	// Create a new Gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req
	c.Set("user", user)

	// Call the GetProfileByID function
	ctrl.GetProfileByID(c)

	// Check the response
	assert.Equal(s.T(), http.StatusOK, record.Code)
	var response map[string]interface{}
	err = json.NewDecoder(record.Body).Decode(&response)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Success", response["message"])
	assert.Equal(s.T(), userExpected, response["data"])
}
