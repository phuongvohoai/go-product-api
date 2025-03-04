package routes

import (
	"phuong/go-product-api/controllers"
	"phuong/go-product-api/database"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(router *gin.RouterGroup) {
	userService := services.NewUserService(database.DB)
	authController := controllers.NewAuthController(*userService)

	router.POST("/auth/login", authController.Login)
}
