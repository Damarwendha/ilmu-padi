package dto

import "github.com/kelompok-2/ilmu-padi/entity"

// Initialize Struct Course ID Dto
type CourseIDDto struct {
	ID     uint   `uri:"id" binding:"required"`
	UserId string `json:"userId,omitempty"`
	User   entity.User
}
