package services

import (
	"context"
	"phuong/go-product-api/database"
	"phuong/go-product-api/models"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (models.Product, error) {
	err := s.db.WithContext(ctx).Create(product).Error
	return *product, err
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (models.Product, error) {
	var product models.Product
	err := s.db.WithContext(ctx).Preload("Category").First(&product, id).Error
	return product, err
}

func (s *ProductService) GetProducts(ctx context.Context, pagination *models.Pagination) ([]models.Product, int, error) {
	var products []models.Product

	query := s.db.WithContext(ctx).Model(&models.Product{}).Preload("Category")

	total, err := database.GetPaginatedList(query, pagination, &products)

	return products, total, err
}

func (s *ProductService) GetProductsByCategory(ctx context.Context, categoryID int, pagination *models.Pagination) ([]models.Product, int, error) {
	var products []models.Product

	query := s.db.WithContext(ctx).Model(&models.Product{}).
		Where("category_id = ?", categoryID).
		Preload("Category")

	total, err := database.GetPaginatedList(query, pagination, &products)

	return products, total, err
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *models.Product) (models.Product, error) {
	var existingProduct models.Product
	err := s.db.WithContext(ctx).First(&existingProduct, product.ID).Error

	if err != nil {
		return models.Product{}, err
	}

	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.CategoryID = product.CategoryID

	err = s.db.WithContext(ctx).Save(&existingProduct).Error
	return existingProduct, err
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	err := s.db.WithContext(ctx).Delete(&models.Product{}, id).Error
	return err
}
