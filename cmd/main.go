package main

import (
	"log"
	"net/http"

	"github.com/CodeChefVIT/cookoff-backend/config"
	"github.com/CodeChefVIT/cookoff-backend/internal/database"
	"github.com/CodeChefVIT/cookoff-backend/internal/models"
	"github.com/CodeChefVIT/cookoff-backend/internal/routes"
	"github.com/CodeChefVIT/cookoff-backend/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load enviorment variables! \n", err.Error())
	}
	database.ConnectDB(&config)
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
	}))

	app.Validator = &utils.CustomValidator{Validator: validator.New()}

	routes.AuthRoutes(app)
	app.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, models.Response{
			Message: "Pong",
			Status:  true,
		})
	})

	app.Logger.Fatal(app.Start(config.Port))
}
