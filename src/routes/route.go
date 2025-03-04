package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		MapPingRoutes(v1)
		MapAuthRoutes(v1)
		MapUserRoutes(v1)
	}

	return router
}
