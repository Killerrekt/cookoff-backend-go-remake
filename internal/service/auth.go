package service

import (
	"github.com/CodeChefVIT/cookoff-backend/internal/database"
	"github.com/CodeChefVIT/cookoff-backend/internal/models"
)

func FindUserByEmail(email string) (models.User, error) {
	var data models.User
	err := database.DB.QueryRow("select * from users where email = $1", email).Scan(
		&data.Name, &data.Email, &data.RegNo, &data.RefreshToken, &data.UserRole,
		&data.IsActive, &data.IsRoundActive, &data.RoundQualified, &data.Password,
		&data.TokenVersion, &data.Score, &data.Submission)
	return data, err
}
