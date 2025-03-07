package services

import (
	"context"
	"errors"
	"phuong/go-product-api/database"
	"phuong/go-product-api/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User, password string) (models.User, error) {
	hashPassword, err := PasswordHash(password)
	if err != nil {
		return models.User{}, err
	}

	user.PasswordHash = hashPassword
	err = s.db.WithContext(ctx).Create(user).Error
	return *user, err
}

func (s *UserService) GetUser(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (s *UserService) GetUsers(ctx context.Context, pagination *models.Pagination) ([]models.User, int, error) {
	var users []models.User

	query := s.db.WithContext(ctx).Model(&models.User{})

	total, err := database.GetPaginatedList(query, pagination, &users)

	return users, total, err
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User, newPassword string) (models.User, error) {
	var existingUser models.User
	err := s.db.WithContext(ctx).First(&existingUser, user.ID).Error

	if err != nil {
		return models.User{}, err
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email

	if newPassword != "" {
		hashPassword, err := PasswordHash(newPassword)
		if err != nil {
			return models.User{}, err
		}
		existingUser.PasswordHash = hashPassword
	}

	s.db.WithContext(ctx).Model(&existingUser).Updates(map[string]interface{}{
		"username":      existingUser.Username,
		"password_hash": existingUser.PasswordHash,
		"email":         existingUser.Email,
	})
	return existingUser, err
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	err := s.db.WithContext(ctx).Delete(&models.User{}, id).Error
	return err
}

func (s *UserService) VerifyUser(ctx context.Context, username, password string) (models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	if !PasswordVerify(password, user.PasswordHash) {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}
