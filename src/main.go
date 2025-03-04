package main

import (
	"phuong/go-product-api/database"
	"phuong/go-product-api/routes"

	"github.com/joho/godotenv"
)

// @title		Product API
// @version		1.0
// @description	API for managing products
func main() {
	godotenv.Load()
	database.Connect()

	router := routes.RegisterRoutes()

	router.Run(":8080")
}
