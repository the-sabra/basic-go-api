package routes

import (
	"firstApi/middlewares"

	"github.com/labstack/echo/v4"
)
  



func SetupRoute(g *echo.Group){

	//Auth/User Route
	auth :=g.Group("/auth")
	AuthRoutes(auth)
	
	
	// after this middleware want to auth
	g.Use(middlewares.JWTMiddleware())
	
	//User Routes
	user  := g.Group("/user")
	UserRoutes(user)
	
	//Book Routes 
	book :=g.Group("/book")
	BookRoutes(book) 
}   