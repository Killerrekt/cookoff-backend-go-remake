package service

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/cookoff-backend/internal/database"
	"github.com/CodeChefVIT/cookoff-backend/internal/models"
)

func FindUserByEmail(email string) (models.User, error) {
	var data models.User
	err := database.DB.QueryRow("select * from users where email = $1", email).Scan(
		&data.Name, &data.Email, &data.RegNo, &data.RefreshToken, &data.UserRole,
		&data.IsActive, &data.IsRoundActive, &data.RoundQualified, &data.Password,
		&data.TokenVersion, &data.Score)
	return data, err
}

func UpdateUserTokenDetails(user models.User) error {
	tx, err := database.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	_, err = tx.Exec("Update users set tokenversion=$1,refreshtoken=$2 where email=$3", user.TokenVersion, user.RefreshToken, user.Email)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

func CreateUser(name string, email string, password string, regno string) error {
	tx, err := database.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into users(name,email,password,regno) values($1,$2,$3,$4)", name, email, password, regno)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
