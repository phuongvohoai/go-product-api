package routes

import (
	"phuong/go-product-api/controllers"
	"phuong/go-product-api/database"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(router *gin.RouterGroup) {
	userService := services.NewUserService(database.DB)
	userController := controllers.NewUserController(*userService)

	userRoute := router.Group("/users")
	{
		userRoute.GET("", userController.GetUsers)
		userRoute.GET("/:id", userController.GetUser)
		userRoute.PUT("/:id", userController.UpdateUser)
		userRoute.DELETE("/:id", userController.DeleteUser)
	}
}
