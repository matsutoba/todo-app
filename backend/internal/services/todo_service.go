package services

import (
	"todo-app/internal/errors"
	"todo-app/internal/models"
	"todo-app/internal/repositories"
)

type ITodoService interface {
	FindAll() (*[]models.Todo, error)
}

type TodoService struct {
	todoRepository repositories.ITodoRepository
}

func NewTodoService(todoRepository repositories.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (t *TodoService) FindAll() (*[]models.Todo, error) {
	todos, err := t.todoRepository.FindAll()
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return &todos, nil
}
