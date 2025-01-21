package services

import (
	"context"
	"todolist-api/internal/models"
	"todolist-api/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

//TODO: Explain why using context
func (s *TodoService) CreateTodo(ctx context.Context, todo *models.CreateTodoInput) (*models.Todo, error) {
	return s.repo.Create(ctx, todo)
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	return s.repo.FindAll(ctx)
}

func (s *TodoService) GetTodoByUserID(ctx context.Context, userId string) ([]*models.Todo, error) {
	return s.repo.FindByUserID(ctx, userId)
}

func (s *TodoService) GetTodoById(ctx context.Context, id string) (*models.Todo, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TodoService) UpdateTodo(ctx context.Context, id string, todo *models.UpdateTodoInput) (*models.Todo, error) {
	return s.repo.Update(ctx, id, todo)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
