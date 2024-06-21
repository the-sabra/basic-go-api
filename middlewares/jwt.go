package middlewares

import (
	"firstApi/util"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := echoJwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")) ,
	} 
	return echoJwt.WithConfig(config) 
} 