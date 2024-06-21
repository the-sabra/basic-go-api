package routes

import (
	"firstApi/handlers"
	"firstApi/middlewares"
	"firstApi/models"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group){
	userRepo := repository.NewGormUserRepo(models.DB)

	userHandler:= handlers.NewUserHandler(userRepo);

	// g.GET("",userHandler.GetAllUsers,middlewares.RoleMiddleware("customer"))
	g.GET("/sigendUser",userHandler.GetUser,middlewares.GetAuthUser()) 
	g.GET("",userHandler.GetAllUsers,middlewares.RoleMiddleware("admin"))
}  