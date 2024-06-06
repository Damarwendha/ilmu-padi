package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

// Initialize Struct User Courses Usecase
type userCoursesFavouriteUsecase struct {
	userCoursesFavouriteRepository repository.UserCoursesFavouriteRepository
}

// Initialize Interface User Course Sender Usecase
type UserCoursesFavouriteUsecase interface {
	AddOrRemoveToFavourite(ucf entity.UserCoursesFavourite) (string, error)
	FindAllByUserID(user_id uint) ([]entity.Course, error)
}

// Construction to Access User Courses Usecase
func NewUserCoursesFavouriteUsecase(userCoursesFavouriteRepository repository.UserCoursesFavouriteRepository) UserCoursesFavouriteUsecase {
	return &userCoursesFavouriteUsecase{
		userCoursesFavouriteRepository: userCoursesFavouriteRepository,
	}
}

// Add Or Remove To Favorite
func (uc *userCoursesFavouriteUsecase) AddOrRemoveToFavourite(ucf entity.UserCoursesFavourite) (string, error) {
	return uc.userCoursesFavouriteRepository.AddOrRemoveToFavourite(ucf)
}

// Get Favorite List
func (uc *userCoursesFavouriteUsecase) FindAllByUserID(user_id uint) ([]entity.Course, error) {
	return uc.userCoursesFavouriteRepository.FindAllByUserID(user_id)
}
