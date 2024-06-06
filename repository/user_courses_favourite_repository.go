package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct User Courses Repository
type userCoursesFavouriteRepository struct {
	db *gorm.DB
}

// Initialize Interface User Courses Sender Repository
type UserCoursesFavouriteRepository interface {
	AddOrRemoveToFavourite(userCourseFavourite entity.UserCoursesFavourite) (string, error)
	FindAllByUserID(userid uint) ([]entity.Course, error)
}

// Construction to Access User Courses Repository
func NewUserCoursesFavouriteRepository(db *gorm.DB) UserCoursesFavouriteRepository {
	return &userCoursesFavouriteRepository{db: db}
}

// Add Or Remove To Favorite
func (ucfr *userCoursesFavouriteRepository) AddOrRemoveToFavourite(userCourseFavourite entity.UserCoursesFavourite) (string, error) {
	var ucf entity.UserCoursesFavourite
	// if its already exits delete, else create
	if err := ucfr.db.Where("user_id = ? AND course_id = ?", userCourseFavourite.UserID, userCourseFavourite.CourseID).First(&ucf).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Add to Favourite Executed", ucfr.db.Create(&userCourseFavourite).Error
		}
		return "", err
	}

	return "Remove from Favourite Executed", ucfr.db.Where("user_id = ? AND course_id = ?", userCourseFavourite.UserID, userCourseFavourite.CourseID).Delete(&entity.UserCoursesFavourite{}).Error
}

// Find All By User ID
func (ucfr *userCoursesFavouriteRepository) FindAllByUserID(userid uint) ([]entity.Course, error) {
	var listRaw []entity.UserCoursesFavourite
	if err := ucfr.db.Where("user_id = ?", userid).Find(&listRaw).Error; err != nil {
		return []entity.Course{}, err
	}

	var list []entity.Course
	for _, userCourseFavourite := range listRaw {
		var course entity.Course
		if err := ucfr.db.Where("id = ?", userCourseFavourite.CourseID).First(&course).Error; err != nil {
			return []entity.Course{}, err
		}

		list = append(list, course)
	}

	return list, nil
}
