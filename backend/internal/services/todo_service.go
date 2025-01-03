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
	FindById(id uint) (*models.Todo, error)
	Update(todoId uint, todo dto.UpdateTodoInput) (*models.Todo, error)
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

func (t *TodoService) FindById(id uint) (*models.Todo, error) {
	todo, err := t.todoRepository.FindById(id)
	if err != nil {
		return nil, errors.ErrTodoNotFound
	}
	return todo, nil
}

func (t *TodoService) Update(todoId uint, updateTodoInput dto.UpdateTodoInput) (*models.Todo, error) {
	dueDate, err := time.Parse("2006-01-02", updateTodoInput.DueDate)
	if err != nil {
		return nil, errors.ErrInvalidRequest
	}

	targetItem, err := t.todoRepository.FindById(todoId)

	if err != nil {
		return nil, errors.ErrTodoNotFound
	}

	targetItem.Title = updateTodoInput.Title
	targetItem.Description = updateTodoInput.Description
	targetItem.Completed = updateTodoInput.Completed
	targetItem.DueDate = dueDate

	todo, err := t.todoRepository.Update(*targetItem)
	if err != nil {
		return nil, errors.ErrUpdateTodoFailed
	}
	return todo, nil
}
