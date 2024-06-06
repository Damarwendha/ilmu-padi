package model

import "github.com/golang-jwt/jwt/v5"

// Struct My Custom Claims
type MyCustomClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`
	Role   string `json:"role"`
}
