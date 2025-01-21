package models

import "time"

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	UserID      string `json:"user_id" validate:"required"`
}

// TODO: Explain why using *
type UpdateTodoInput struct {
	Title       string `json:"title" validate:"omitempty,required_without_all=Description Completed"`
	Description string `json:"description" validate:"omitempty,required_without_all=Title Completed"`
	Completed   bool   `json:"completed" validate:"omitempty,required_without_all=Title Description"`
}

