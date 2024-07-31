package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !isValidToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Dummy token validation function (replace with actual logic)
func isValidToken(token string) bool {
	// Example token validation logic (replace with actual logic)
	return strings.HasPrefix(token, "Bearer ")
}
