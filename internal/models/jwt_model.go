package models

import "github.com/golang-jwt/jwt/v5"

type AccessToken struct {
	RegNo        string
	Role         string
	TokenVersion int
	jwt.RegisteredClaims
}

type RefreshToken struct {
	RegNo string
	Role  string
	jwt.RegisteredClaims
}
