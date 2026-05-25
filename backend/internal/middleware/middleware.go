package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["id"])
		c.Set("role_level", claims["role_level"])

		c.Next()
	}
}
