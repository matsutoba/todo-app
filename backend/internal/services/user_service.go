package services

import (
	"todo-app/internal/dto"
	"todo-app/internal/models"
	"todo-app/internal/repositories"
)

type IUserService interface {
	Create(newUser dto.CreateUserInput) (*models.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) Create(createUserInput dto.CreateUserInput) (*models.User, error) {
	newUser := models.User{
		Email:        createUserInput.Email,
		PasswordHash: createUserInput.Passowrd,
	}
	return u.userRepository.Create(newUser)
}
