package repositories

import (
	"todo-app/internal/errors"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(todo models.Todo) (*models.Todo, error)
	FindAll() ([]models.Todo, error)
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo models.Todo) (*models.Todo, error) {
	result := r.db.Create(&todo)
	if result.Error != nil {
		return nil, errors.ErrInsertFailed
	}
	return &todo, nil
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, errors.ErrNotFound
	}
	return todos, nil
}
