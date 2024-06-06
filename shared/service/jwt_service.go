package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

type JwtService interface {
	CreateToken(author entity.User) (dto.AuthResponseDto, error)
	ParseToken(tokenString string) (map[string]interface{}, error)
}

type jwtService struct {
	cfg config.TokenConfig
}

func (j *jwtService) CreateToken(user entity.User) (dto.AuthResponseDto, error) {
	claims := model.MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.TokenIssue,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.TokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: strconv.FormatUint(uint64(user.ID), 10),
		Role:   user.Role,
	}

	token := jwt.NewWithClaims(j.cfg.SigningMethod, claims)
	ss, err := token.SignedString(j.cfg.TokenSecret)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("oops, failed to create token: %v", err)
	}
	return dto.AuthResponseDto{Token: ss}, nil
}

func (s *jwtService) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.cfg.TokenSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("oops, failed to parse token claims")
	}
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}
