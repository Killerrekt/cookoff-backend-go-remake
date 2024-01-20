package utils

import (
	"time"

	"github.com/CodeChefVIT/cookoff-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func CreateAccessToken(user models.User, exp time.Duration) (string, error) {
	claims := &models.AccessToken{
		RegNo:            user.RegNo,
		Role:             user.UserRole,
		TokenVersion:     user.TokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("ACCESS_SECRET_KEY")))
}

func CreateRefreshToken(user models.User, exp time.Duration) (string, error) {
	claims := &models.RefreshToken{
		RegNo:            user.RegNo,
		Role:             user.UserRole,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("REFRESH_SECRET_KEY")))
}
