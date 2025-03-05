package routes

import (
	docs "phuong/go-product-api/docs"
	"phuong/go-product-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		authorized := v1.Group("")
		{
			authorized.Use(middleware.AuthenticateMiddleware())
			MapUserRoutes(authorized)
			MapCategoryRoutes(authorized)
			MapProductRoutes(authorized)
		}

		// Public routes without authentication
		MapPingRoutes(v1)
		MapAuthRoutes(v1)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
