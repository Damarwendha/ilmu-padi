package dto

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

// Initialize Struct Register Dto
type RegisterDto struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Initialize Struct Login Dto
type LoginDto struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

// Initialize Struct Logout Dto
type LogoutDto struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

// Initialize Struct Auth Response Dto
type AuthResponseDto struct {
	Token string `json:"token"`
}

// Initialize Struct Verification Token Dto
type VerificationTokenDto struct {
	gorm.Model
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Initialize Struct User Response Dto
type UserResponseDto struct {
	gorm.Model
	Name                string                `json:"name"`
	Email               string                `json:"email,omitempty"`
	Password            string                `json:"password"`
	Role                string                `json:"role"`
	Subscription_Status string                `json:"subscription_status"`
	Courses             []entity.Course       `json:"courses"`
	Subscriptions       []entity.Subscription `json:"subscriptions"`
}

// Initialize Struct Forgot Password Dto
type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required,email"`
}

// Initialize Struct Reset Password Dto
type ResetPasswordDto struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
