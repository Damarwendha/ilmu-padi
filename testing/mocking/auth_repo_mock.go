package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/mock"
)

type AuthRepoMock struct {
	mock.Mock
}

func (m *AuthRepoMock) FindByEmailAuth(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *AuthRepoMock) FindByResetToken(resetToken string) (*entity.User, error) {
	args := m.Called(resetToken)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *AuthRepoMock) Save(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *AuthRepoMock) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *AuthRepoMock) FindByVerificationToken(token string) (*entity.User, error) {
	args := m.Called(token)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *AuthRepoMock) UpdateResetToken(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *AuthRepoMock) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}
