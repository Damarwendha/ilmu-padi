package usecase

import (
	"testing"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	arm *mocking.UserRepoMock
	auc UserUsecase
}

func (s *UserUsecaseTestSuite) SetupTest() {
	s.arm = new(mocking.UserRepoMock)
	s.auc = NewUserUsecase(s.arm)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (s *UserUsecaseTestSuite) TestFindAll_Success() {
	// Setup
	expected := []entity.User{}
	s.arm.On("FindAll").Return(expected, model.Paging{}, nil)

	// Execute
	users, _, err := s.auc.FindAll(1, 10, "")

	// Assert
	s.NoError(err)
	s.Equal(expected, users)
	s.arm.AssertExpectations(s.T())
}

func (s *UserUsecaseTestSuite) TestGetUserByID_Success() {
	// Mocking
	userID := uint(1)
	mockUser := entity.User{
		ID: int(userID),
	}

	// Expected Result
	expectedUser := mockUser

	// Setup
	s.arm.On("GetProfileByID", userID).Return(mockUser, nil)

	// Action
	user, err := s.auc.GetProfileByID(userID, "")

	// Assertion
	s.NoError(err)
	s.Equal(expectedUser, user)

	// Verify
	s.arm.AssertExpectations(s.T())
}
