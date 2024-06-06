package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/mock"
)

type PaymentRepoMock struct {
	mock.Mock
}

func (m *PaymentRepoMock) FindAll() ([]entity.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]entity.Transaction), args.Error(2)
}

func (m *PaymentRepoMock) GetByCourseID(courseID int) ([]entity.Transaction, error) {
	args := m.Called(courseID)
	return args.Get(0).([]entity.Transaction), args.Error(1)
}

func (m *PaymentRepoMock) GetByID(id int) (entity.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *PaymentRepoMock) Save(transaction entity.Transaction) (entity.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *PaymentRepoMock) Update(transaction entity.Transaction) (entity.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(entity.Transaction), args.Error(1)
}
