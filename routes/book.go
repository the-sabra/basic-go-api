package routes

import (
	"firstApi/handlers"
	"firstApi/middlewares"
	"firstApi/models"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func BookRoutes(g *echo.Group){

	bookRepo := repository.NewGormBookRepo(models.DB)

	bookHandler:= handlers.NewBookHandler(bookRepo); 
	 
	g.POST("", bookHandler.CreateBook, middlewares.GetAuthUser())
	g.GET("",bookHandler.GetBooks)
	g.GET("/:id", bookHandler.GetBook)
	g.PATCH("/:id",bookHandler.UpdateBook, middlewares.GetAuthUser()) 
	g.DELETE("/:id",bookHandler.DeleteBook)
}  