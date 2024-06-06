package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}
	s.repo = NewUserRepository(gormDb)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestFindAll_Success() {
	const (
		page  = 1
		size  = 10
		total = 100
	)

	// Test case: Page exists and size is less than total rows
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users LIMIT 10 OFFSET 0")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	users, paging, err := s.repo.FindAll(page, size)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), page, paging.Page)
	assert.Equal(s.T(), size, paging.RowsPerPage)
	assert.Equal(s.T(), total, paging.TotalRows)
	assert.Equal(s.T(), 10, paging.TotalPages)
	assert.NotEmpty(s.T(), users)

	// Test case: Page does not exist
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users LIMIT 10 OFFSET 10")).
		WillReturnError(sql.ErrNoRows)

	users, _, err = s.repo.FindAll(2, size)
	assert.Error(s.T(), err)
	assert.Empty(s.T(), users)
}

func (s *UserRepositoryTestSuite) TestFindByID_Success() {
	// Test case: User exists
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE id = ? LIMIT 1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	user, err := s.repo.FindByID(1)
	s.NoError(err)
	s.NotEmpty(user)

	// Test case: User does not exist
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE id = ? LIMIT 1")).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)

	user, err = s.repo.FindByID(2)
	s.Error(err)
	s.Empty(user)
}

func (s *UserRepositoryTestSuite) TestGetRolesByUserID_Success() {
	// Test case: User exists
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users JOIN user_roles on roles.id = user_roles.role_id WHERE user_roles.user_id = ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	roles, err := s.repo.GetRolesByUserID(1)
	s.NoError(err)
	s.NotEmpty(roles)

	// Test case: User does not exist
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users JOIN user_roles on roles.id = user_roles.role_id WHERE user_roles.user_id = ?")).
		WithArgs(2).
		WillReturnError(sql.ErrNoRows)

	roles, err = s.repo.GetRolesByUserID(2)
	s.Error(err)
	s.Empty(roles)
}

func (s *UserRepositoryTestSuite) TestFindByEmailUser_Success() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE email = ? LIMIT 1")).
		WithArgs("john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
			AddRow(1, "John Doe", "john@example.com", "user"))

	user, err := s.repo.FindByEmailUser("john@example.com")
	s.NoError(err)
	s.NotEmpty(user)
}

// Test case: User does not exist
func (s *UserRepositoryTestSuite) TestFindByEmailUser_NotFound() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM users WHERE email = ? LIMIT 1")).
		WithArgs("jane@example.com").
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := s.repo.FindByEmailUser("jane@example.com")
	s.Error(err)
	s.Empty(user)
	s.Equal(gorm.ErrRecordNotFound, err)
}
