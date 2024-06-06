package entity

import (
	"time"

	"gorm.io/gorm"
)

// Initialize Struct Subscription
type Subscription struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	User_ID    string    `json:"user_id"`
	Start_Date time.Time `json:"start_date"`
	End_Date   time.Time `json:"end_date"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
