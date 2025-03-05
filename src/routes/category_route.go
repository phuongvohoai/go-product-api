package routes

import (
	"phuong/go-product-api/controllers"
	"phuong/go-product-api/database"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func MapCategoryRoutes(router *gin.RouterGroup) {
	categoryService := services.NewCategoryService(database.DB)
	categoryController := controllers.NewCategoryController(categoryService)
	categoryRoute := router.Group("/categories")
	{
		categoryRoute.GET("", categoryController.GetCategories)
		categoryRoute.GET("/:id", categoryController.GetCategoryById)
		categoryRoute.POST("", categoryController.CreateCategory)
		categoryRoute.PUT("/:id", categoryController.UpdateCategory)
		categoryRoute.DELETE("/:id", categoryController.DeleteCategory)
	}
}
