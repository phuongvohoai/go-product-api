package routes

import (
	docs "phuong/go-product-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		MapPingRoutes(v1)
		MapAuthRoutes(v1)
		MapUserRoutes(v1)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
