package routes

import (
	"phuong/go-product-api/controllers"

	"github.com/gin-gonic/gin"
)

func MapPingRoutes(router *gin.RouterGroup) {
	pingController := controllers.NewPingController()

	router.GET("/ping", pingController.Ping)
}
