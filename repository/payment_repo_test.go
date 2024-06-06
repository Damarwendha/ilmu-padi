package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type PaymentRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    PaymentRepository
}

func (s *PaymentRepositoryTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}
	s.repo = NewPaymentRepository(gormDb)
}

func TestPaymentRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentRepositoryTestSuite))
}

func (s *PaymentRepositoryTestSuite) TestGetByCourseID_Success() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "transactions" WHERE course_id = $1 ORDER BY id desc`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "course_id", "user_id", "amount"}).
			AddRow(1, 1, 1, 100).
			AddRow(2, 1, 2, 200))

	transactions, err := s.repo.GetByCourseID(1)
	s.NoError(err)
	s.Len(transactions, 2)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *PaymentRepositoryTestSuite) TestGetByCourseID_Error() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "transactions" WHERE course_id = $1 ORDER BY id desc`)).
		WithArgs(1).
		WillReturnError(gorm.ErrRecordNotFound)

	transactions, err := s.repo.GetByCourseID(1)
	s.Error(err)
	s.Empty(transactions)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *PaymentRepositoryTestSuite) TestGetByID_Success() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "transactions" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "course_id", "user_id", "amount"}).
			AddRow(1, 1, 1, 100))

	transaction, err := s.repo.GetByID(1)
	s.NoError(err)
	s.NotEmpty(transaction)
	s.Equal(1, transaction.ID)
	s.Equal(1, transaction.Course_ID)
	s.Equal(1, transaction.User_ID)
	s.Equal(100, transaction.Amount)
	s.NoError(s.mockSql.ExpectationsWereMet())
}

func (s *PaymentRepositoryTestSuite) TestGetByID_Error() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "transactions" WHERE id = $1`)).
		WithArgs(1).
		WillReturnError(gorm.ErrRecordNotFound)

	transaction, err := s.repo.GetByID(1)
	s.Error(err)
	s.Empty(transaction)
	s.NoError(s.mockSql.ExpectationsWereMet())
}
