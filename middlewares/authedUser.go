package middlewares

import (
	"firstApi/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetAuthUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*util.JwtCustomClaims)

			if(claims.Role == "") {
				return echo.ErrUnauthorized
			}

			c.Set("user_id",claims.UserId)
			c.Set("role", claims.Role)
			return next(c); 
		}
	}
} 	  