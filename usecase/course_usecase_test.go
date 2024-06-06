package usecase

import (
	"testing"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/suite"
)

type CourseUsecaseTestSuite struct {
	suite.Suite
	arm *mocking.CourseRepoMock
	auc CourseUsecase
}

func (s *CourseUsecaseTestSuite) SetupTest() {
	s.arm = new(mocking.CourseRepoMock)
	s.auc = NewCourseUsecase(s.arm)
}

func (s *CourseUsecaseTestSuite) TestGetAllCourse_Success() {
	s.arm.On("FindAll").Return([]entity.Course{}, model.Paging{}, nil)
	courses, _, err := s.auc.GetAllCourses(1, 10, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal([]entity.Course{}, courses)
}

func (s *CourseUsecaseTestSuite) TestGetAllCourseByCategory_Success() {
	s.arm.On("FindAllByCategory", "").Return([]entity.Course{}, model.Paging{}, nil)
	courses, _, err := s.auc.GetAllCoursesByCategory("", 1, 10, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal([]entity.Course{}, courses)
}

func (s *CourseUsecaseTestSuite) TestGetCourseByID_Success() {
	s.arm.On("FindByID", 1).Return(entity.Course{}, nil)
	course, err := s.auc.GetCourseByID(1, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal(entity.Course{}, course)
}

func (s *CourseUsecaseTestSuite) TestCreateCourse_Success() {
	s.arm.On("Create", entity.Course{}).Return(nil)
	course, err := s.auc.CreateCourse(entity.Course{}, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal(entity.Course{}, course)
}

func (s *CourseUsecaseTestSuite) TestUpdateCourse_Success() {
	courseID := 1
	course := entity.Course{
		ID:              int(courseID),
		Title:           "New Title",
		Description:     "New Description",
		Category:        "New Category",
		Video_URL:       "New Video URL",
		Duration:        10,
		Instructor_Name: "New Instructor",
		Rating:          5,
	}

	s.arm.On("FindByID", courseID).Return(course, nil)
	s.arm.On("Update", course).Return(nil)

	updatedCourse, err := s.auc.UpdateCourse(courseID, course, "")
	s.Nil(err)
	s.Equal(course, updatedCourse)

	s.arm.AssertExpectations(s.T())
}

func (s *CourseUsecaseTestSuite) TestDeleteCourse_Success() {
	s.arm.On("Delete", 1).Return(nil)
	err := s.auc.DeleteCourse(1, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
}

func TestCourseUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseUsecaseTestSuite))
}
