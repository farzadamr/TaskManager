package services

import (
	"context"
	"errors"

	"github.com/farzadamr/TaskManager/models"
	"github.com/farzadamr/TaskManager/repositories"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, userID uint, name string) (*models.Category, error)
	GetUserCategories(ctx context.Context, userID uint) ([]*models.Category, error)
	GetCategoryByID(ctx context.Context, userID, categoryID uint) (*models.Category, error)
	DeleteCategory(ctx context.Context, userID, categoryID uint) error
}
type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) CreateCategory(ctx context.Context, userID uint, name string) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("category name connot be empty")
	}

	category := &models.Category{
		Name:   name,
		UserID: userID,
	}

	if err := s.categoryRepo.Create(ctx, *category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetUserCategories(ctx context.Context, userID uint) ([]*models.Category, error) {
	return s.categoryRepo.FindByUserID(ctx, userID)
}

func (s *categoryService) GetCategoryByID(ctx context.Context, userID, categoryID uint) (*models.Category, error) {
	return s.categoryRepo.FindByIDAndUserID(ctx, userID, categoryID)
}
func (s *categoryService) DeleteCategory(ctx context.Context, userID, categoryID uint) error {
	return s.categoryRepo.Delete(ctx, categoryID, userID)
}
