package routes

import (
	"firstApi/handlers"
	"firstApi/middlewares"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func BookRoutes(g *echo.Group) {
	bookRepo := repository.NewGormBookRepo(repository.DB)

	bookHandler := handlers.NewBookHandler(bookRepo)

	g.POST("", bookHandler.CreateBook, middlewares.GetAuthUser())
	g.GET("", bookHandler.GetBooks)
	g.GET("/:id", bookHandler.GetBook)
	g.PATCH("/:id", bookHandler.UpdateBook, middlewares.GetAuthUser())
	g.DELETE("/:id", bookHandler.DeleteBook)
}
