package routes

import (
	"firstApi/middlewares"
	"firstApi/repository"

	"github.com/labstack/echo/v4"
)

func SetupRoute(g *echo.Group, repo repository.Repository) {
	// Auth/User Route
	auth := g.Group("/auth")
	AuthRoutes(auth, repo)

	// after this middleware want to auth
	g.Use(middlewares.JWTMiddleware())

	// User Routes
	user := g.Group("/user")
	UserRoutes(user, repo)

	// Book Routes
	book := g.Group("/book")
	BookRoutes(book, repo)
}

