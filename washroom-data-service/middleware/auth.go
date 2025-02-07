package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware provides authentication for protected routes
// Currently checks for presence of Authorization header
// TODO: Implement proper token validation
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// TODO: Implement actual token validation
		// For now, just checking if header exists

		c.Next()
	}
}
