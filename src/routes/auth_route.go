package routes

import (
	"phuong/go-product-api/controllers"
	"phuong/go-product-api/database"
	"phuong/go-product-api/middleware"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(router *gin.RouterGroup) {
	userService := services.NewUserService(database.DB)
	emailService := services.NewLocalMailService()
	authController := controllers.NewAuthController(*userService, emailService)
	userController := controllers.NewUserController(*userService)

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", userController.CreateUser)

	router.POST("/auth/logout", middleware.AuthenticateMiddleware(), authController.Logout)
}
