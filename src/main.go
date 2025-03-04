package main

import (
	"phuong/go-product-api/database"
	"phuong/go-product-api/routes"
)

func main() {
	database.Connect()

	router := routes.RegisterRoutes()

	router.Run(":8080")
}
