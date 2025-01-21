package repository

import (
	"context"
	"fmt"
	"todolist-api/internal/models"
	"todolist-api/prisma/db"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.CreateUserInput) (*models.User, error)
	Update(ctx context.Context, id string, user *models.UpdateUserInput) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	client *db.PrismaClient
}

func NewUserRepository(client *db.PrismaClient) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, user *models.CreateUserInput) (*models.User, error) {
	createdUser, err := r.client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
		db.User.Username.Set(user.Username),
	).Exec(ctx)

  if err != nil {
		return nil, err
	}
	fmt.Println(createdUser)

	return &models.User{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (r *userRepository) Update(ctx context.Context, id string, user *models.UpdateUserInput) (*models.User, error) {
  updatedUser, err := r.client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Update(
    db.User.Username.Set(user.Username),
    db.User.Email.Set(user.Email),
    db.User.Password.Set(user.Password),
  ).Exec(ctx)

  if err != nil {
    return nil, err
  }

  return &models.User{
    ID:        updatedUser.ID,
    Username:  updatedUser.Username,
    Email:     updatedUser.Email,
    CreatedAt: updatedUser.CreatedAt,
    UpdatedAt: updatedUser.UpdatedAt,
  }, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*models.User, error) {
  users, err := r.client.User.FindMany().Exec(ctx)

  if err != nil {
    return nil, err
  }

  var result []*models.User

  for _, user := range users {
    result = append(result, &models.User{
      ID:        user.ID,
      Username:  user.Username,
      Email:     user.Email,
      CreatedAt: user.CreatedAt,
      UpdatedAt: user.UpdatedAt,
    })
  }

  return result, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
  user, err := r.client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Exec(ctx)

  if err != nil {
    return nil, err
  }

  return &models.User{
    ID:        user.ID,
    Username:  user.Username,
    Email:     user.Email,
    CreatedAt: user.CreatedAt,
    UpdatedAt: user.UpdatedAt,
  }, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
  user, err := r.client.User.FindUnique(
    db.User.Email.Equals(email),
  ).Exec(ctx)

  if err != nil {
    return nil, err
  }

  return &models.User{
    ID:        user.ID,
    Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
