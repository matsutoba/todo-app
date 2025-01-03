package services

import (
	"time"
	"todo-app/internal/dto"
	"todo-app/internal/errors"
	"todo-app/internal/models"
	"todo-app/internal/repositories"
)

type ITodoService interface {
	Create(todo dto.CreateTodoInput) (*models.Todo, error)
	FindAll() (*[]models.Todo, error)
}

type TodoService struct {
	todoRepository repositories.ITodoRepository
}

func NewTodoService(todoRepository repositories.ITodoRepository) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (t *TodoService) Create(createTodoInput dto.CreateTodoInput) (*models.Todo, error) {
	dueDate, err := time.Parse("2006-01-02", createTodoInput.DueDate)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}

	newTodo := models.Todo{
		Title:       createTodoInput.Title,
		Description: createTodoInput.Description,
		Completed:   createTodoInput.Completed,
		DueDate:     dueDate,
	}

	todo, err := t.todoRepository.Create(newTodo)
	if err != nil {
		return nil, errors.ErrCreateTodoFailed
	}
	return todo, nil
}

func (t *TodoService) FindAll() (*[]models.Todo, error) {
	todos, err := t.todoRepository.FindAll()
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return &todos, nil
}
