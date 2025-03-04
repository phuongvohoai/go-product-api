package database

import (
	"log"

	"phuong/go-product-api/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "sqlserver://sa:Password1@@localhost:1433?database=ProductApi&connection+timeout=30"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to SQLServer:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln("Cannot migrate table:", err)
	}

	DB = db
}
