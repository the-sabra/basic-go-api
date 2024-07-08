package routes

import (
	"firstApi/handlers"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group, userRepo repository.UserRepository) {
	userHandler := handlers.NewUserHandler(userRepo)

	g.POST("/signup", userHandler.SignUp)
	g.POST("/login", userHandler.Login)
}
