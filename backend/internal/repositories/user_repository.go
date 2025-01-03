package repositories

import (
	"todo-app/internal/errors"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(newUser models.User) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return nil, errors.ErrDuplicateEntry
	}
	return &newUser, nil
}
