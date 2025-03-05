package middleware

import (
	"phuong/go-product-api/models"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			c.AbortWithStatusJSON(400, models.Response.BadRequest(err))
		}
	}
}
