package middleware

import (
	"github.com/CodeChefVIT/cookoff-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AccessTokenProtected(app *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.AccessToken)
		},
		SigningKey: []byte(viper.GetString("ACCESS_KEY_SECRET")),
	}
	app.Use(echojwt.WithConfig(config))
}

func RefreshTokenProtected(app *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.RefreshToken)
		},
		SigningKey: []byte(viper.GetString("REFRESH_KEY_SECRET")),
	}
	app.Use(echojwt.WithConfig(config))
}
