package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *mocking.CourseUsecaseMock
	amm *mocking.AuthMiddlewareMock
}

func (s *CourseControllerTestSuite) SetupTest() {
	s.amm = new(mocking.AuthMiddlewareMock)
	s.aum = new(mocking.CourseUsecaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")
	rg.Use(s.amm.RequireToken("user", "admin"))
	s.rg = rg
}

func (s *CourseControllerTestSuite) TestGetAllCourse_Success() {
	s.aum.On("GetAllCourses", 1, 10, "1").Return([]entity.Course{}, model.Paging{}, nil)

	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/courses", nil)
	assert.NoError(s.T(), err)
	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetAllCourses(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestGetAllCourseCategory_Success() {
	s.aum.On("GetAllCoursesByCategory", "tech", 1, 10, "1").Return([]entity.Course{}, model.Paging{}, nil)

	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/courses/category/tech", nil)
	assert.NoError(s.T(), err)
	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetAllCoursesByCategory(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestGetAllCourse_Failed() {
	s.aum.On("GetAllCourses", 1, 10, "1").Return([]entity.Course{}, model.Paging{}, fmt.Errorf("Get All Course Failed"))

	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/courses", nil)
	assert.NotNil(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetAllCourses(c)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestGetCourse_Success() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetCourseByID", "1").Return(1, "1")
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/course/1", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetCourseByID(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestGetCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetCourse", "1").Return("Deuhh", "hiohf")
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/course/foeufo", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetCourseByID(c)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestCreateCourse_Success() {
	payload := entity.Course{
		ID:              1,
		Title:           "Test",
		Description:     "Test",
		Category:        "Test",
		Video_URL:       "Test",
		Duration:        0,
		Instructor_Name: "0",
		Rating:          0,
		UserId:          "0",
	}

	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("CreateCourse", payload, "1").Return(payload, nil)

	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/api/v1/asset/courses", bytes.NewBuffer(jsonPayload))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.CreateCourse(c)

	assert.Equal(s.T(), http.StatusCreated, record.Code)
}

func (s *CourseControllerTestSuite) TestCreateCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("CreateCourse", entity.Course{}).Return(fmt.Errorf("Create Course Failed"))
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("POST", "/api/v1/asset/courses", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.CreateCourse(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_Success() {
	// Mocking the required methods
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("UpdateCourse", 1, entity.Course{UserId: "1"}, "1").Return(entity.Course{}, nil)

	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", bytes.NewBufferString(`{}`))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.WithValue(req.Context(), "user", "1"))

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Calling the UpdateCourse method
	coursesCtrl.UpdateCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("UpdateCourse", 1, entity.Course{}, "1").Return(entity.Course{}, fmt.Errorf("Update Course Failed"))
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", nil)
	assert.NotNil(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.UpdateCourse(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func TestCourseControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CourseControllerTestSuite))
}

func (s *CourseControllerTestSuite) TestUpdateCourse_InvalidID() {
	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request with invalid ID
	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/invalid", nil)
	assert.NoError(s.T(), err)

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Calling the UpdateCourse method
	coursesCtrl.UpdateCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_BindUriError() {
	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", bytes.NewBufferString(`{}`))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Mocking the binding error
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	c.Set("user", "1")

	// Calling the UpdateCourse method
	coursesCtrl.UpdateCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_BindJSONError() {
	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request with invalid JSON
	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", bytes.NewBufferString(`invalid_json`))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Mocking the binding error
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	c.Set("user", "1")

	// Calling the UpdateCourse method
	coursesCtrl.UpdateCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_UpdateError() {
	// Mocking the required methods
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("UpdateCourse", 1, entity.Course{UserId: "1"}, "1").Return(entity.Course{}, errors.New("update error"))

	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", bytes.NewBufferString(`{}`))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(context.WithValue(req.Context(), "user", "1"))

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Calling the UpdateCourse method
	coursesCtrl.UpdateCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}
func (s *CourseControllerTestSuite) TestDeleteCourse_Success() {
	// Mocking the required methods
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("DeleteCourse", 1, "1").Return(nil)

	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("DELETE", "/api/v1/asset/courses/1", nil)
	assert.NoError(s.T(), err)
	req = req.WithContext(context.WithValue(req.Context(), "user", "1"))

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	// Calling the DeleteCourse method
	coursesCtrl.DeleteCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestDeleteCourse_InvalidID() {
	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request with invalid ID
	req, err := http.NewRequest("DELETE", "/api/v1/asset/courses/invalid", nil)
	assert.NoError(s.T(), err)

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Calling the DeleteCourse method
	coursesCtrl.DeleteCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestDeleteCourse_BindUriError() {
	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("DELETE", "/api/v1/asset/courses/1", nil)
	assert.NoError(s.T(), err)
	req = req.WithContext(context.WithValue(req.Context(), "user", "1"))

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	// Mocking the binding error
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	c.Set("user", "1")

	// Calling the DeleteCourse method
	coursesCtrl.DeleteCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *CourseControllerTestSuite) TestDeleteCourse_DeleteError() {
	// Mocking the required methods
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("DeleteCourse", 1, "1").Return(errors.New("delete error"))

	// Initializing the CourseController
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Creating the HTTP request
	req, err := http.NewRequest("DELETE", "/api/v1/asset/courses/1", nil)
	assert.NoError(s.T(), err)
	req = req.WithContext(context.WithValue(req.Context(), "user", "1"))

	// Recording the HTTP response
	record := httptest.NewRecorder()

	// Setting up the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	// Calling the DeleteCourse method
	coursesCtrl.DeleteCourse(c)

	// Asserting the response status code
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *CourseControllerTestSuite) TestCreateCourse() {
	// Setup
	coursesCtrl := NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	// Create the HTTP request
	payload := entity.Course{Title: "Test Course", Description: "Test Description"}
	json, _ := common.ToJSON(payload)
	req, err := http.NewRequest("POST", "/api/v1/asset/courses", json)
	s.NoError(err)

	// Record the HTTP response
	record := httptest.NewRecorder()

	// Setup the gin context
	c, _ := gin.CreateTestContext(record)
	c.Request = req
	c.Set("user", "testuser")

	// Call the CreateCourse function
	coursesCtrl.CreateCourse(c)

	// Assert the response
	s.Equal(http.StatusOK, record.Code)
	resp := entity.Course{}
	s.NotNil(resp)

	// Assert the course in the database
	course, err := s.aum.FindByID(int(resp.ID))
	s.NoError(err)
	s.Equal("Test Course", course.Title)
	s.Equal("Test Description", course.Description)
}
