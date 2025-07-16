package services

import (
	"context"
	"errors"

	"github.com/farzadamr/TaskManager/models"
	"github.com/farzadamr/TaskManager/repositories"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*models.Task, error)
	GetTask(ctx context.Context, id uint) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	UpdateTask(ctx context.Context, id uint, title, description string, completed bool) (*models.Task, error)
	DeleteTask(ctx context.Context, id uint) error
	CompleteTask(ctx context.Context, id uint) (*models.Task, error)
}

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string) (*models.Task, error) {
	if title == "" {
		return nil, errors.New("title connot be empty")
	}
	task := &models.Task{
		Title:       title,
		Description: description,
	}
	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetTask(ctx context.Context, id uint) (*models.Task, error) {
	return s.taskRepo.FindByID(ctx, id)
}
func (s *taskService) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	return s.taskRepo.FindAll(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, id uint, title, description string, completed bool) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if title != "" {
		task.Title = title
	}

	if description != "" {
		task.Description = description
	}

	task.Completed = completed

	if err := s.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(ctx context.Context, id uint) error {
	return s.taskRepo.Delete(ctx, id)
}

func (s *taskService) CompleteTask(ctx context.Context, id uint) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if task.Completed {
		return task, nil
	}

	if err := s.taskRepo.MarkAsComplete(ctx, id); err != nil {
		return nil, err
	}
	task.Completed = true
	return task, nil
}
