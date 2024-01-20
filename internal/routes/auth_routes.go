package routes

import (
	"github.com/CodeChefVIT/cookoff-backend/internal/controller"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(app *echo.Echo) {
	auth := app.Group("/auth")
	auth.POST("/signup", controller.SignUp)
	auth.POST("/login", controller.Login)
}
