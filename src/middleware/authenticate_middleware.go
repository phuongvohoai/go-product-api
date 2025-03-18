package middleware

import (
	"errors"
	"phuong/go-product-api/models"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationValue := c.GetHeader("Authorization")
		if authorizationValue == "" {
			c.AbortWithStatusJSON(400, models.Response.BadRequest(errors.New("MISSING Authorization Header")))
			return
		}

		bearerToken := authorizationValue[7:] // Remove "Bearer " prefix
		validToken, claims := services.VerifyToken(bearerToken)
		if !validToken {
			c.AbortWithStatusJSON(401, models.Response.BadRequest(errors.New("INVALID TOKEN")))
			return
		}

		if services.IsTokenRevoked(claims.StandardClaims.Id) {
			c.AbortWithStatusJSON(401, models.Response.BadRequest(errors.New("TOKEN IS REVOKED")))
			return
		}

		// Store claims in request context
		c.Set("Claims", claims)

		c.Next()
	}
}
