package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Initialize Struct User Usecase
type userUsecase struct {
	userRepo repository.UserRepository
}

// Initialize Interface User Sender Usecase
type UserUsecase interface {
	FindAll(page, size int, user string) ([]entity.User, model.Paging, error)
	GetProfileByID(userID uint, user string) (entity.User, error)
}

// Construction to Access User Usecase
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

// Find All
func (u *userUsecase) FindAll(page, size int, user string) ([]entity.User, model.Paging, error) {
	return u.userRepo.FindAll(page, size)
}

// Get Profile By ID
func (u *userUsecase) GetProfileByID(userID uint, user string) (entity.User, error) {
	return u.userRepo.FindByID(userID)
}
