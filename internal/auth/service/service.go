package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/auth/models"
	"github.com/imirjar/Michman/internal/auth/storage"
)

type Config interface {
	GetSecret() string
}

type Storage interface {
	GetUserID(ctx context.Context, username string) (int, error)
}

type Service struct {
	config  Config
	storage Storage
}

func NewService() *Service {
	return &Service{
		config:  config.NewConfig(),
		storage: storage.NewStorage(),
	}
}

func (s *Service) BuildJWTString(ctx context.Context, user models.User) (string, error) {

	userID, err := s.storage.GetUserID(ctx, user.Username)
	if err != nil {
		return err.Error(), err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(s.config.GetSecret()))
	if err != nil {
		return err.Error(), err
	}

	return tokenString, nil
}

func (s *Service) GetUserID(ctx context.Context, token string) int {
	claims := &models.Claims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.config.GetSecret()), nil
	})
	return claims.UserID
}

func (s *Service) CreateNewUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}
