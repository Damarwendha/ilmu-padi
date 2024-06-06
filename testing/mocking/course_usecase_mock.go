package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/stretchr/testify/mock"
)

type CourseUsecaseMock struct {
	mock.Mock
}

// CreateCourse implements usecase.CourseUsecase.
func (m *CourseUsecaseMock) CreateCourse(courses entity.Course, user string) (entity.Course, error) {
	args := m.Called(courses, user)
	return args.Get(0).(entity.Course), args.Error(1)
}

func (m *CourseUsecaseMock) GetAllCoursesByCategory(category string, page int, size int, user string) ([]entity.Course, model.Paging, error) {
	args := m.Called(category, page, size, user)
	return args.Get(0).([]entity.Course), args.Get(1).(model.Paging), args.Error(2)
}

// DeleteCourse implements usecase.CourseUsecase.
func (m *CourseUsecaseMock) DeleteCourse(id int, user string) error {
	args := m.Called(id, user)
	return args.Error(0)
}

// GetAllCourses implements usecase.CourseUsecase.
func (m *CourseUsecaseMock) GetAllCourses(page int, size int, user string) ([]entity.Course, model.Paging, error) {
	args := m.Called(page, size, user)
	return args.Get(0).([]entity.Course), args.Get(1).(model.Paging), args.Error(2)
}

// GetCourseByID implements usecase.CourseUsecase.
func (m *CourseUsecaseMock) GetCourseByID(id int, user string) (entity.Course, error) {
	args := m.Called(id, user)
	return args.Get(0).(entity.Course), args.Error(1)
}

// UpdateCourse implements usecase.CourseUsecase.
func (m *CourseUsecaseMock) UpdateCourse(id int, courses entity.Course, user string) (entity.Course, error) {
	args := m.Called(id, courses, user)
	return args.Get(0).(entity.Course), args.Error(1)
}

func (m *CourseUsecaseMock) FindAll() ([]entity.Course, error) {
	args := m.Called()
	return args.Get(0).([]entity.Course), args.Error(1)
}

func (m *CourseUsecaseMock) FindByID(id int) (entity.Course, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Course), args.Error(1)
}
