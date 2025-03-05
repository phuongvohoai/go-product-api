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

func GetPaginatedList(query *gorm.DB, pagination *models.Pagination, dest interface{}) (int, error) {
	var total int64

	countQuery := query.Session(&gorm.Session{})
	if err := countQuery.Count(&total).Error; err != nil {
		return 0, err
	}

	if err := query.
		Order(pagination.SortBy + " " + pagination.SortDir).
		Limit(pagination.PageSize).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Find(dest).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}
