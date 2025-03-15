package database

import (
	"log"
	"os"
	"time"

	"phuong/go-product-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	maxRetries := 5

	db, err := connectWithRetry(dsn, maxRetries)
	if err != nil {
		log.Fatalln("Maximum connection attempts reached. Cannot connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})
	if err != nil {
		log.Fatalln("Cannot migrate table:", err)
	}

	DB = db
}

func connectWithRetry(dsn string, maxRetries int) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	retryCount := 0

	for retryCount < maxRetries {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}

		log.Printf("Cannot connect to database (attempt %d/%d): %v", retryCount+1, maxRetries, err)
		retryCount++

		if retryCount < maxRetries {
			retryDelay := time.Duration(retryCount) * time.Second
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}
	return nil, err
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
