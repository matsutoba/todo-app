package repositories

import (
	"todo-app/internal/errors"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	FindAll() ([]models.Todo, error)
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, errors.ErrNotFound
	}
	return todos, nil
}
