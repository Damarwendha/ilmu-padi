package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) FindAll(page, size int, user string) ([]entity.User, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.User), args.Get(1).(model.Paging), args.Error(2)
}

func (m *UserUsecaseMock) GetProfileByID(id uint, user string) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}
