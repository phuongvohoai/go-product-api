package database

import (
	"log"
	"os"

	"phuong/go-product-api/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to SQLServer:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
	if err != nil {
		log.Fatalln("Cannot migrate table:", err)
	}

	DB = db
}
