package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserCoursesFavourite struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string `json:"user_id"`
	CourseID  string `json:"course_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
