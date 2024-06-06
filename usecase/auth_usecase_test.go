package usecase

import (
	"errors"
	"testing"

	"github.com/kelompok-2/ilmu-padi/client"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthUsecaseTestSuite struct {
	suite.Suite
	arm *mocking.AuthRepoMock
	ajm *mocking.JwtServMock
	amm *mocking.MailServMock
	auc AuthUsecase
}

func (s *AuthUsecaseTestSuite) SetupTest() {
	s.arm = new(mocking.AuthRepoMock)
	s.auc = NewAuthUsecase(s.arm, s.ajm, s.amm)
}

func TestAuthUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}

func (s *AuthUsecaseTestSuite) TestRegister_Success() {
	hashedPassword := "hashedPassword"
	verificationToken := "verificationToken"

	mockRegisterDto := &dto.RegisterDto{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
		Role:     "user",
	}

	configMock := new(mock.Mock)
	configMock.On("HashPassword", mockRegisterDto.Password).Return(hashedPassword, nil)

	utilsMock := new(mock.Mock)
	utilsMock.On("FormatTemplate", client.VerifyEmailTemplate, mock.Anything).Return("formattedHTML")

	s.arm.On("Create", mock.Anything).Return(nil)
	s.arm.On("Update", mock.Anything).Return(nil)
	s.arm.On("GenerateVerificationToken").Return(verificationToken, nil)

	user, err := s.auc.Register(mockRegisterDto)
	s.NoError(err)
	s.NotEmpty(user)
	s.Equal(user.Name, mockRegisterDto.Name)
	s.Equal(user.Email, mockRegisterDto.Email)
	s.Equal(user.Password, hashedPassword)
	s.Equal(user.Role, mockRegisterDto.Role)
	s.Equal(user.VerificationToken, verificationToken)

	s.arm.AssertExpectations(s.T())
}

func (s *AuthUsecaseTestSuite) TestRegister_HashPasswordError() {
	mockRegisterDto := &dto.RegisterDto{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
		Role:     "user",
	}

	configMock := new(mock.Mock)
	configMock.On("HashPassword", mockRegisterDto.Password).Return("", assert.AnError)

	user, err := s.auc.Register(mockRegisterDto)
	s.Error(err)
	s.Empty(user)

	s.arm.AssertNotCalled(s.T(), "Create", mock.Anything)
	s.arm.AssertNotCalled(s.T(), "Update", mock.Anything)
}

func (s *AuthUsecaseTestSuite) TestLogin_Success() {
	email := "test@example.com"
	password := "password"

	mockLoginDto := &dto.LoginDto{
		Email:    email,
		Password: password,
	}

	mockUser := &entity.User{
		Email:    email,
		Password: "hashedPassword",
		Verified: true,
	}

	mockTokenDetails := &dto.AuthResponseDto{
		Token: "token",
	}

	s.arm.On("FindByEmailAuth", email).Return(mockUser, nil)
	s.ajm.On("CreateToken", *mockUser).Return(mockTokenDetails, nil)

	token, err := s.auc.Login(mockLoginDto)
	s.NoError(err)
	s.Equal(mockTokenDetails.Token, token)

	s.arm.AssertExpectations(s.T())
	s.ajm.AssertExpectations(s.T())
}

func (s *AuthUsecaseTestSuite) TestLogin_UserNotVerified() {
	email := "test@example.com"
	password := "password"

	mockLoginDto := &dto.LoginDto{
		Email:    email,
		Password: password,
	}

	mockUser := &entity.User{
		Email:    email,
		Password: "hashedPassword",
		Verified: false,
	}

	s.arm.On("FindByEmailAuth", email).Return(mockUser, nil)

	token, err := s.auc.Login(mockLoginDto)
	s.Error(err)
	s.Equal("", token)

	s.arm.AssertExpectations(s.T())
}

func (s *AuthUsecaseTestSuite) TestLogin_InvalidEmailOrPassword() {
	email := "test@example.com"
	password := "password"

	mockLoginDto := &dto.LoginDto{
		Email:    email,
		Password: password,
	}

	mockUser := &entity.User{
		Email:    email,
		Password: "hashedPassword",
		Verified: true,
	}

	s.arm.On("FindByEmailAuth", email).Return(mockUser, nil)

	token, err := s.auc.Login(mockLoginDto)
	s.Error(err)
	s.Equal("", token)

	s.arm.AssertExpectations(s.T())
}

func (s *AuthUsecaseTestSuite) TestLogin_FindByEmailAuthError() {
	email := "test@example.com"
	password := "password"

	mockLoginDto := &dto.LoginDto{
		Email:    email,
		Password: password,
	}

	mockError := errors.New("some error")

	s.arm.On("FindByEmailAuth", email).Return(nil, mockError)

	token, err := s.auc.Login(mockLoginDto)
	s.Error(err)
	s.Equal("", token)

	s.arm.AssertExpectations(s.T())
}
