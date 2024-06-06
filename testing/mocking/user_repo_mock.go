package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) FindByID(id uint) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserRepoMock) FindAll(page, size int) ([]entity.User, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.User), args.Get(1).(model.Paging), args.Error(2)
}

func (m *UserRepoMock) FindByEmailUser(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *UserRepoMock) GetRolesByUserID(userID uint) ([]entity.User, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.User), args.Error(1)
}
