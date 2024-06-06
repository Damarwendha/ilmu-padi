package entity

import (
	"time"

	"gorm.io/gorm"
)

// Initialize Struct Course
type Course struct {
	ID              int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Video_URL       string  `json:"video_url"`
	Duration        int     `json:"duration"`
	Instructor_Name string  `json:"instructor_name"`
	Rating          float32 `json:"rating"`
	UserId          string  `json:"userId,omitempty"`
	BackerCount     int
	GoalAmount      int
	CurrentAmount   int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
