package middleware

import (
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationValue := c.GetHeader("Authorization")
		if authorizationValue == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		bearerToken := authorizationValue[7:] // Remove "Bearer " prefix
		validToken := services.VerifyToken(bearerToken)
		if !validToken {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
