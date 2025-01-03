package repositories

import (
	"todo-app/internal/errors"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(todo models.Todo) (*models.Todo, error)
	FindAll() ([]models.Todo, error)
	FindById(id uint) (*models.Todo, error)
	Update(todo models.Todo) (*models.Todo, error)
	Delete(id uint) error
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

func (r *TodoRepository) FindById(id uint) (*models.Todo, error) {
	var todo models.Todo
	result := r.db.First(&todo, id)
	if result.Error != nil {
		return nil, errors.ErrNotFound
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo models.Todo) (*models.Todo, error) {
	result := r.db.Save(&todo)
	if result.Error != nil {
		return nil, errors.ErrUpdateFailed
	}
	return &todo, nil
}

func (r *TodoRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return errors.ErrNotFound
	}
	return nil
}
