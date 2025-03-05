package main

import (
	"phuong/go-product-api/database"
	"phuong/go-product-api/middleware"
	"phuong/go-product-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title		Product API
// @version		1.0
// @description	API for managing products
func main() {
	godotenv.Load()
	database.Connect()

	r := gin.Default()

	r.Use(middleware.ErrorHandlingMiddleware())
	r.Use(middleware.PerformanceLogMiddleware())
	r.Use(middleware.ResponseCacheMiddleware())

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
