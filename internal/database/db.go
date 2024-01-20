package database

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/CodeChefVIT/cookoff-backend/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDB(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = sql.Open("pgx", dsn)
	fmt.Println(dsn)

	if err != nil {
		log.Fatalln("Failed to connect to the Database! \n", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")
}
