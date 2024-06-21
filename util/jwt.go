package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

  
type JwtCustomClaims struct{
	UserId uint 
	Role string
	jwt.RegisteredClaims
	}	
	
	func GenerateJWT(claims JwtCustomClaims) string{
	claims.ExpiresAt =  jwt.NewNumericDate(time.Now().Add(time.Minute * 55))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	signedToken ,err :=token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{ 
		return ""
	}

	return signedToken  
}    