package repository

import (
	"context"
	"todolist-api/internal/models"
	"todolist-api/prisma/db"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *models.CreateTodoInput) (*models.Todo, error)
	FindAll(ctx context.Context) ([]*models.Todo, error)
	FindByID(ctx context.Context, id string) (*models.Todo, error)
	FindByUserID(ctx context.Context, userId string) ([]*models.Todo, error)
	Update(ctx context.Context, id string, todo *models.UpdateTodoInput) (*models.Todo, error)
	Delete(ctx context.Context, id string) error
}

type todoRepository struct {
	client *db.PrismaClient
}

func NewTodoRepository(client *db.PrismaClient) TodoRepository {
	return &todoRepository{client: client}
}

func (r *todoRepository) Create(ctx context.Context, todo *models.CreateTodoInput) (*models.Todo, error) {
	createdTodo, err := r.client.Todo.CreateOne(
		db.Todo.Title.Set(todo.Title),
		db.Todo.Description.Set(todo.Description),
		db.Todo.User.Link(
			db.User.ID.Equals(todo.UserID),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return r.mapToTodo(createdTodo), nil
}

func (r *todoRepository) FindAll(ctx context.Context) ([]*models.Todo, error) {
	todos, err := r.client.Todo.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	var todoList []*models.Todo
	for _, todo := range todos {
		todoList = append(todoList, r.mapToTodo(&todo))
	}

	return todoList, nil
}

func (r *todoRepository) FindByUserID(ctx context.Context, userId string) ([]*models.Todo, error) {
	todos, err := r.client.Todo.FindMany(
		db.Todo.User.Where(
			db.User.ID.Equals(userId),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var todoList []*models.Todo
	for _, todo := range todos {
		todoList = append(todoList, r.mapToTodo(&todo))
	}

	return todoList, nil
}

func (r *todoRepository) FindByID(ctx context.Context, id string) (*models.Todo, error) {
	todo, err := r.client.Todo.FindUnique(
		db.Todo.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return r.mapToTodo(todo), nil
}

func (r *todoRepository) Update(ctx context.Context, id string, todo *models.UpdateTodoInput) (*models.Todo, error) {
	updatedTodo, err := r.client.Todo.FindUnique(db.Todo.ID.Equals(id)).Update(
		db.Todo.Title.Set(todo.Title),
		db.Todo.Description.Set(todo.Description),
		db.Todo.Completed.Set(todo.Completed),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return r.mapToTodo(updatedTodo), nil
}

func (r *todoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.client.Todo.FindUnique(db.Todo.ID.Equals(id)).Delete().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) mapToTodo(todo *db.TodoModel) *models.Todo {
	return &models.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
