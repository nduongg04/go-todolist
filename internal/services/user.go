package services

import (
	"context"
	"todolist-api/internal/models"
	"todolist-api/internal/repository"
	"todolist-api/pkg/errors"
	"todolist-api/pkg/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.CreateUserInput) (*models.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &models.UserResponse{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *UserService) Login(ctx context.Context, input *models.LoginUserInput) (*models.UserResponse, error) {
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.UserNotFound
	}

	if !utils.ComparePassword(user.Password, input.Password) {
		return nil, errors.InvalidCredentials
	}

	return &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
