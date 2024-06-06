package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Payment Repository
type paymentRepository struct {
	db *gorm.DB
}

// Initialize Interface Payment Sender Repository
type PaymentRepository interface {
	GetByCourseID(courseID int) ([]entity.Transaction, error)
	GetByUserID(userID int) ([]entity.Transaction, error)
	GetByID(ID int) (entity.Transaction, error)
	Save(transaction entity.Transaction) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	FindAll() ([]entity.Transaction, error)
}

// Construction to Access Payment Repository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// Save Payment
func (r *paymentRepository) GetByCourseID(courseID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("users").Preload("courses").Where("course_id = ?", courseID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *paymentRepository) GetByUserID(userID int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("users").Preload("courses").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *paymentRepository) GetByID(ID int) (entity.Transaction, error) {
	var transaction entity.Transaction

	err := r.db.Preload("users").Preload("courses").Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *paymentRepository) Save(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *paymentRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *paymentRepository) FindAll() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := r.db.Preload("users").Preload("courses").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
