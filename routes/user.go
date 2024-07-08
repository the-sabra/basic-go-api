package routes

import (
	"firstApi/handlers"
	"firstApi/middlewares"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group, userRepo repository.UserRepository) {
	userHandler := handlers.NewUserHandler(userRepo)

	// g.GET("",userHandler.GetAllUsers,middlewares.RoleMiddleware("customer"))
	g.GET("/sigendUser", userHandler.GetUser, middlewares.GetAuthUser())
	g.GET("", userHandler.GetAllUsers, middlewares.RoleMiddleware("admin"))
}
