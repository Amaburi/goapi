package middleware

import (
	"goapi/utils" // Replace with your actual project path
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

func isValidToken(tokenStr string) bool {
	// Remove "Bearer " prefix
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// Parse and validate the token
	_, err := utils.ParseToken(tokenStr)
	return err == nil
}
