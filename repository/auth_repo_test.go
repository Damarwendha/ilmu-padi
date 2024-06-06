package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/suite"
)

type AuthRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    AuthRepository
}

func (s *AuthRepositoryTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}
	s.repo = NewAuthRepository(gormDb)
}

func TestAuthRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepositoryTestSuite))
}

func (s *AuthRepositoryTestSuite) TestFindByResetToken_Success() {
	// Test case: User exists
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE reset_token = ? LIMIT 1")).
		WithArgs("testToken").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	user, err := s.repo.FindByResetToken("testToken")
	s.NoError(err)
	s.NotEmpty(user)

	// Test case: User does not exist
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE reset_token = ? LIMIT 1")).
		WithArgs("notFoundToken").
		WillReturnError(gorm.ErrRecordNotFound)

	user, err = s.repo.FindByResetToken("notFoundToken")
	s.Error(err)
	s.Empty(user)
}

func (s *AuthRepositoryTestSuite) TestFindByEmailAuth_Success() {
	// Test case: User exists
	query := regexp.QuoteMeta("SELECT * FROM users WHERE email = ? LIMIT 1")
	s.mockSql.ExpectQuery(query).
		WithArgs("john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	user, err := s.repo.FindByEmailAuth("john@example.com")
	s.NoError(err)
	s.NotEmpty(user)

	// Test case: User does not exist
	s.mockSql.ExpectQuery(query).
		WithArgs("johnnotfound@example.com").
		WillReturnError(gorm.ErrRecordNotFound)

	user, err = s.repo.FindByEmailAuth("johnnotfound@example.com")
	s.Error(err)
	s.Empty(user)
}

func (s *AuthRepositoryTestSuite) TestSave_Success() {
	user := &entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "user",
	}

	s.mockSql.ExpectBegin()
	s.mockSql.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","role") VALUES ($1,$2,$3)`)).
		WithArgs(user.Name, user.Email, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockSql.ExpectCommit()

	err := s.repo.Save(user)
	s.NoError(err)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *AuthRepositoryTestSuite) TestUpdateResetToken_Success() {
	user := &entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "user",
	}

	s.mockSql.ExpectBegin()
	s.mockSql.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","role") VALUES ($1,$2,$3)`)).
		WithArgs(user.Name, user.Email, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockSql.ExpectCommit()

	err := s.repo.UpdateResetToken(user)
	s.NoError(err)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *AuthRepositoryTestSuite) TestUpdate_Success() {
	user := &entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "user",
	}

	s.mockSql.ExpectBegin()
	s.mockSql.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","role") VALUES ($1,$2,$3)`)).
		WithArgs(user.Name, user.Email, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockSql.ExpectCommit()

	err := s.repo.Update(user)
	s.NoError(err)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *AuthRepositoryTestSuite) TestCreate_Success() {
	user := &entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "user",
	}

	s.mockSql.ExpectBegin()
	s.mockSql.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","role") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs(user.Name, user.Email, user.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mockSql.ExpectCommit()

	err := s.repo.Create(user)
	s.NoError(err)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *AuthRepositoryTestSuite) TestCreate_Error() {
	user := &entity.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "user",
	}

	s.mockSql.ExpectBegin()
	s.mockSql.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","email","role") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs(user.Name, user.Email, user.Role).
		WillReturnError(gorm.ErrInvalidTransaction)
	s.mockSql.ExpectRollback()

	err := s.repo.Create(user)
	s.Error(err)
	s.NoError(s.mockSql.ExpectationsWereMet())
}
