package repositories

import (
	"context"

	"github.com/farzadamr/TaskManager/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	FindByID(ctx context.Context, id uint) (*models.Task, error)
	FindAll(ctx context.Context) ([]*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uint) error
	MarkAsComplete(ctx context.Context, id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) FindByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) FindAll(ctx context.Context) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *taskRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, id).Error
}

func (r *taskRepository) MarkAsComplete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Task{}).Where("id = ?", id).Update("completed", true).Error
}
