package models

import "time"

type User struct {
	Name           string
	Email          string
	RegNo          string
	RefreshToken   string
	IsActive       bool
	IsRoundActive  bool
	RoundQualified int
	Password       string
	UserRole       string
	TokenVersion   int
	Score          int
	Submission     time.Time
}
