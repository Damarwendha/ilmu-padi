package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/stretchr/testify/mock"
)

type CourseRepoMock struct {
	mock.Mock
}

func (m *CourseRepoMock) FindAll(page, size int) ([]entity.Course, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.Course), args.Get(1).(model.Paging), args.Error(2)
}

func (m *CourseRepoMock) FindAllByCategory(category string, page, size int) ([]entity.Course, model.Paging, error) {
	args := m.Called(category, page, size)
	return args.Get(0).([]entity.Course), args.Get(1).(model.Paging), args.Error(2)
}

func (m *CourseRepoMock) FindByID(id int) (entity.Course, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Course), args.Error(1)
}

func (m *CourseRepoMock) Create(course *entity.Course) error {
	args := m.Called(course)
	return args.Error(0)
}

func (m *CourseRepoMock) Update(course entity.Course) error {
	args := m.Called(&course)
	return args.Error(0)
}

func (m *CourseRepoMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
