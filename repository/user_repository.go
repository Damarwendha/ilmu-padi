package repository

import (
	"log"
	"math"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Initialize Struct User Repository
type userRepository struct {
	db *gorm.DB
}

// Initialize Interface User Sender Repository
type UserRepository interface {
	FindAll(page, size int) ([]entity.User, model.Paging, error)
	FindByID(id uint) (entity.User, error)
	GetRolesByUserID(userID uint) ([]entity.User, error)
	FindByEmailUser(email string) (entity.User, error)
}

// Construction to Access User Repository
func NewUserRepository(db *gorm.DB) UserRepository {
	if db == nil {
		log.Fatal("Database connection is nil UserRepository")
	}

	return &userRepository{db: db}
}

// Find All
func (u *userRepository) FindAll(page, size int) ([]entity.User, model.Paging, error) {
	if u.db == nil {
		log.Fatal("Database connection is nil in FindAll")
	}

	var users []entity.User
	offset := (page - 1) * size

	// Calculate the row total first
	var totalRows int
	if err := u.db.Model(&entity.User{}).Count(&totalRows).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Retrieve data with limits and offsets for pagination
	if err := u.db.Limit(size).Offset(offset).Find(&users).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Set up paging information
	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return users, paging, nil
}

// Find By ID
func (u *userRepository) FindByID(id uint) (entity.User, error) {
	if u.db == nil {
		log.Fatal("Database connection is nil in FindByID")
	}

	var user entity.User
	if err := u.db.First(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Get Roles By User ID
func (u *userRepository) GetRolesByUserID(userID uint) ([]entity.User, error) {
	if u.db == nil {
		log.Fatal("Database connection is nil in GetRolesByUserID")
	}

	var roles []entity.User
	if err := u.db.Joins("JOIN user_roles on roles.id = user_roles.role_id").Where("user_roles.user_id = ?", userID).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Find By Email User
func (u *userRepository) FindByEmailUser(email string) (entity.User, error) {
	if u.db == nil {
		log.Fatal("Database connection is nil in FindByEmailUser")
	}
	var user entity.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
