package utils

import (
	"time"

	"diflow/api/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, username string) (string, error) {
	cfg := config.Load()
	expirationTime := time.Now().Add(time.Duration(cfg.JWT.AccessExpiry) * time.Minute)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	cfg := config.Load()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateRefreshToken(userID uint, username string) (string, error) {
	cfg := config.Load()
	expirationTime := time.Now().Add(time.Duration(cfg.JWT.RefreshExpiry) * 24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

