package repositories

import (
	"context"

	"github.com/farzadamr/TaskManager/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, category models.Category) error
	FindByUserID(ctx context.Context, userID uint) ([]models.Category, error)
	FindByIDAndUserID(ctx context.Context, userID, categoryID uint) (*models.Category, error)
	Delete(ctx context.Context, categoryID, userID uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Create(ctx context.Context, category models.Category) error {
	return c.db.WithContext(ctx).Create(category).Error
}

func (c *categoryRepository) FindByUserID(ctx context.Context, userID uint) ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.WithContext(ctx).Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryRepository) FindByIDAndUserID(ctx context.Context, userID, categoryID uint) (*models.Category, error) {
	var category *models.Category
	if err := c.db.WithContext(ctx).Where("user_id = ?", userID).Find(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) Delete(ctx context.Context, categoryID, userID uint) error {
	return c.db.WithContext(ctx).Where("id = ? AND user_id = ?", categoryID, userID).Delete(&models.Category{}).Error
}
