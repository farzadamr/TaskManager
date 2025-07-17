package services

import (
	"context"

	"github.com/farzadamr/TaskManager/models"
	"github.com/farzadamr/TaskManager/repositories"
	"github.com/farzadamr/TaskManager/utils"
)

type AuthService interface {
	Register(ctx context.Context, username, name, email, password string) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(ctx context.Context, username, name, email, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", nil
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", nil
	}

	return token, nil
}
