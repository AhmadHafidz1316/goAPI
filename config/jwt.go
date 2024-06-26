package config

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("jokowigaming123")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Tidak ada token"})
			return
		}

		// Split the header to get the token part
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		claims := &JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid Token"})
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}