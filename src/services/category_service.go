package services

import (
	"context"
	"phuong/go-product-api/models"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) (models.Category, error) {
	err := s.db.WithContext(ctx).Create(category).Error
	return *category, err
}

func (s *CategoryService) GetCategory(ctx context.Context, id int) (models.Category, error) {
	var category models.Category
	err := s.db.WithContext(ctx).First(&category, id).Error
	return category, err
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := s.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category *models.Category) (models.Category, error) {
	var existingCategory models.Category
	err := s.db.WithContext(ctx).First(&existingCategory, category.ID).Error

	if err != nil {
		return models.Category{}, err
	}

	existingCategory.Name = category.Name
	existingCategory.Description = category.Description

	err = s.db.WithContext(ctx).Save(&existingCategory).Error
	return existingCategory, err
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {
	// First check if there are any products using this category
	var count int64
	s.db.WithContext(ctx).Model(&models.Product{}).Where("category_id = ?", id).Count(&count)
	if count > 0 {
		return gorm.ErrForeignKeyViolated
	}

	err := s.db.WithContext(ctx).Delete(&models.Category{}, id).Error
	return err
}
