package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Auth Repository
type authRepository struct {
	db *gorm.DB
}

// Initialize Interface Auth Sender Repository
type AuthRepository interface {
	FindByResetToken(token string) (*entity.User, error)
	FindByVerificationToken(token string) (*entity.User, error)
	FindByEmailAuth(email string) (*entity.User, error)
	Save(user *entity.User) error
	Create(user *entity.User) error
	UpdateResetToken(user *entity.User) error
	Update(user *entity.User) error
}

// Construction to Access Auth Repository
func NewAuthRepository(db *gorm.DB) AuthRepository {
	if db == nil {
		log.Fatal("Database connection is nil AuthRepository")
	}

	return &authRepository{db: db}
}

// Find By Reset Token
func (a *authRepository) FindByResetToken(token string) (*entity.User, error) {
	var user entity.User
	if err := a.db.Where("reset_token = ?", token).First(&user).Error; err != nil {
		return &entity.User{}, err
	}
	return &user, nil
}

// Find by Verification Token
func (a *authRepository) FindByVerificationToken(token string) (*entity.User, error) {
	var user entity.User
	if err := a.db.Where("verification_token = ?", token).First(&user).Error; err != nil {
		return &entity.User{}, err
	}
	return &user, nil
}

// Find By Email Auth
func (a *authRepository) FindByEmailAuth(email string) (*entity.User, error) {
	var user entity.User
	if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
		return &entity.User{}, err
	}
	return &user, nil
}

// Save
func (a *authRepository) Save(user *entity.User) error {
	return a.db.Save(&user).Error
}

// Create
func (a *authRepository) Create(user *entity.User) error {
	log.Printf("Creating user: %+v\n", user)

	user.ID = 0
	err := a.db.Create(user).Error
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
	}
	return err
}

// Update Reset Token
func (a *authRepository) UpdateResetToken(user *entity.User) error {
	return a.db.Save(user).Error
}

// Update
func (a *authRepository) Update(user *entity.User) error {
	return a.db.Save(user).Error
}
