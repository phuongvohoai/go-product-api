package routes

import (
	"phuong/go-product-api/controllers"
	"phuong/go-product-api/database"
	"phuong/go-product-api/services"

	"github.com/gin-gonic/gin"
)

func MapProductRoutes(router *gin.RouterGroup) {
	productService := services.NewProductService(database.DB)
	productController := controllers.NewProductController(productService)
	productRoute := router.Group("/products")
	{
		productRoute.GET("", productController.GetProducts)
		productRoute.GET("/:id", productController.GetProductById)
		productRoute.POST("", productController.CreateProduct)
		productRoute.PUT("/:id", productController.UpdateProduct)
		productRoute.DELETE("/:id", productController.DeleteProduct)
	}
}
