package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/mock"
)

type CourseFavouriteRepoMock struct {
	mock.Mock
}

func (m *CourseFavouriteRepoMock) AddOrRemoveToFavourite(userCourseFavourite entity.UserCoursesFavourite) (string, error) {
	args := m.Called(userCourseFavourite)
	return args.String(0), args.Error(1)
}

func (m *CourseFavouriteRepoMock) FindAllByUserID(userid uint) ([]entity.Course, error) {
	args := m.Called(userid)
	return args.Get(0).([]entity.Course), args.Error(1)
}
