package repository

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseFavouriteRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserCoursesFavouriteRepository
}

var payload = entity.UserCoursesFavourite{
	ID:       1,
	UserID:   "1",
	CourseID: "1",
}

func (s *CourseFavouriteRepoTestSuite) TestAddOrRemoveCourseFavourite_Added() {
	sqlmock.NewRows([]string{"id", "user_id", "course_id"}).AddRow(1, payload.UserID, payload.CourseID)

	s.mockSql.ExpectQuery(`SELECT \* FROM "user_courses_favourites" WHERE "user_courses_favourites"."deleted_at" IS NULL AND \(\(user_id = \$1 AND course_id = \$2\)\) ORDER BY "user_courses_favourites"."id" ASC LIMIT 1`).
		WithArgs(payload.UserID, payload.CourseID).WillReturnError(gorm.ErrRecordNotFound)

	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`INSERT INTO user_courses_favourites (user_id, course_id) VALUES ($1, $2)`)).
		WithArgs(payload.UserID, payload.CourseID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mockSql.ExpectCommit()

	result, err := s.repo.AddOrRemoveToFavourite(payload)

	s.NoError(err)
	s.Equal("Add to Favourite Executed", result)

	err = s.mockSql.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CourseFavouriteRepoTestSuite) TestAddOrRemoveCourseFavourite_Removed() {
	s.mockSql.ExpectQuery(`SELECT \* FROM "user_courses_favourites" WHERE "user_courses_favourites"."deleted_at" IS NULL AND \(\(user_id = \$1 AND course_id = \$2\)\) ORDER BY "user_courses_favourites"."id" ASC LIMIT 1`).
		WithArgs(payload.UserID, payload.CourseID).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "course_id"}).AddRow(payload.UserID, payload.CourseID))

	s.mockSql.ExpectBegin()

	s.mockSql.ExpectExec(regexp.QuoteMeta(`UPDATE "user_courses_favourites" SET "deleted_at"=$1 WHERE "user_courses_favourites"."deleted_at" IS NULL AND ((user_id = $2 AND course_id = $3))`)).
		WithArgs(time.Now(), payload.UserID, payload.CourseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mockSql.ExpectCommit()

	result, err := s.repo.AddOrRemoveToFavourite(payload)

	s.NoError(err)
	s.Equal("Removed from Favourite Executed", result)

	err = s.mockSql.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CourseFavouriteRepoTestSuite) TestFindAllByUserID() {
	expectedCoursesFav := []entity.UserCoursesFavourite{
		{
			ID:       1,
			UserID:   "1",
			CourseID: "1",
		},
		{
			ID:       2,
			UserID:   "1",
			CourseID: "2",
		},
	}
	expectedCourses := []entity.Course{
		{
			ID:              1,
			Title:           "title",
			Description:     "description",
			UserId:          "1",
			Category:        "category",
			Video_URL:       "video_url",
			Duration:        0,
			Instructor_Name: "instructor_name",
			Rating:          1,
		},
		{
			ID:              2,
			Title:           "title",
			Description:     "description",
			UserId:          "1",
			Category:        "category",
			Video_URL:       "video_url",
			Duration:        0,
			Instructor_Name: "instructor_name",
			Rating:          1,
		}}

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user_courses_favourites" WHERE "user_courses_favourites"."deleted_at" IS NULL AND ((user_id = $1))`)).
		WithArgs(uint(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "course_id"}).AddRow(expectedCoursesFav[0].ID, expectedCoursesFav[0].UserID, expectedCoursesFav[0].CourseID).AddRow(expectedCoursesFav[1].ID, expectedCoursesFav[1].UserID, expectedCoursesFav[1].CourseID))

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE "courses"."deleted_at" IS NULL AND ((id = $1)) ORDER BY "courses"."id" ASC LIMIT 1`)).
		WithArgs(expectedCoursesFav[0].CourseID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "user_id", "category", "video_url", "duration", "instructor_name", "rating", "created_at", "updated_at"}).AddRow(expectedCourses[0].ID, expectedCourses[0].Title, expectedCourses[0].Description, expectedCourses[0].UserId, expectedCourses[0].Category, expectedCourses[0].Video_URL, expectedCourses[0].Duration, expectedCourses[0].Instructor_Name, expectedCourses[0].Rating, expectedCourses[0].CreatedAt, expectedCourses[0].UpdatedAt))

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "courses" WHERE "courses"."deleted_at" IS NULL AND ((id = $1)) ORDER BY "courses"."id" ASC LIMIT 1`)).
		WithArgs(expectedCoursesFav[1].CourseID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "user_id", "category", "video_url", "duration", "instructor_name", "rating", "created_at", "updated_at"}).AddRow(expectedCourses[1].ID, expectedCourses[1].Title, expectedCourses[1].Description, expectedCourses[1].UserId, expectedCourses[1].Category, expectedCourses[1].Video_URL, expectedCourses[1].Duration, expectedCourses[1].Instructor_Name, expectedCourses[1].Rating, expectedCourses[1].CreatedAt, expectedCourses[1].UpdatedAt))

	result, err := s.repo.FindAllByUserID(uint(1))

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)

	err = s.mockSql.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *CourseFavouriteRepoTestSuite) SetupTest() {
	s.mockDb, s.mockSql, _ = sqlmock.New()
	gormDb, err := gorm.Open("postgres", s.mockDb)
	if err != nil {
		panic(err)
	}

	s.repo = NewUserCoursesFavouriteRepository(gormDb)
}

func TestCourseFavouriteRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CourseFavouriteRepoTestSuite))
}
