package repository

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Initialize Struct Course Repository
type courseRepository struct {
	db *gorm.DB
}

// Initialize Interface Course Sender Repository
type CourseRepository interface {
	Create(course *entity.Course) error
	FindAll(page, size int) ([]entity.Course, model.Paging, error)
	FindAllByCategory(category string, page, size int) ([]entity.Course, model.Paging, error)
	FindByID(ID int) (entity.Course, error)
	Update(course entity.Course) error
	Delete(ID int) error
}

// Construction to Access Course Repository
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

// Create
func (c *courseRepository) Create(course *entity.Course) error {
	return c.db.Create(&course).Error
}

// Find All
func (c *courseRepository) FindAll(page, size int) ([]entity.Course, model.Paging, error) {
	var courses []entity.Course
	offset := (page - 1) * size

	// Calculate the row total first
	var totalRows int
	if err := c.db.Model(&entity.Course{}).Count(&totalRows).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Retrieve data with limits and offsets for pagination
	if err := c.db.Limit(size).Offset(offset).Find(&courses).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Set up paging information
	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return courses, paging, nil
}

func (c *courseRepository) FindAllByCategory(category string, page, size int) ([]entity.Course, model.Paging, error) {
	var courses []entity.Course
	offset := (page - 1) * size

	// Calculate the row total first
	var totalRows int
	if err := c.db.Model(&entity.Course{}).Count(&totalRows).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Retrieve data with limits and offsets for pagination
	if err := c.db.Limit(size).Offset(offset).Where("category = ?", category).Find(&courses).Error; err != nil {
		return nil, model.Paging{}, err
	}

	// Set up paging information
	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return courses, paging, nil
}

// Find By ID
func (c *courseRepository) FindByID(id int) (entity.Course, error) {
	var course entity.Course
	if err := c.db.First(&course, id).Error; err != nil {
		return course, err
	}
	return course, nil
}

// Update
func (c *courseRepository) Update(course entity.Course) error {
	return c.db.Save(&course).Error
}

// Delete
func (c *courseRepository) Delete(id int) error {
	return c.db.Delete(&entity.Course{}, id).Error
}
