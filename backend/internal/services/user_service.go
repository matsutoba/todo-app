package services

import (
	"todo-app/internal/dto"
	"todo-app/internal/models"
	"todo-app/internal/repositories"
	"todo-app/internal/utils"
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
	hashedPassword, err := utils.HashPassword(createUserInput.Passowrd)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Email:        createUserInput.Email,
		PasswordHash: hashedPassword,
	}
	return u.userRepository.Create(newUser)
}
