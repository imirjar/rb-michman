package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}
