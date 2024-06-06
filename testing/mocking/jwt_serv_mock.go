package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/stretchr/testify/mock"
)

type JwtServMock struct {
	mock.Mock
}

func (j *JwtServMock) CreateToken(user entity.User) (dto.AuthResponseDto, error) {
	args := j.Called(user)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (j *JwtServMock) ParseToken(token string) (map[string]interface{}, error) {
	args := j.Called(token)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}
