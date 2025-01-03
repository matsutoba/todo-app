package services

import (
	stdErrors "errors"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/dto"
	"todo-app/internal/errors"
	"todo-app/internal/models"
	"todo-app/internal/repositories"
	"todo-app/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

type IUserService interface {
	Create(newUser dto.CreateUserInput) (*models.User, error)
	Login(user dto.LoginUserInput) (*string, error)
	GetUserFromToken(token string) (*models.User, error)
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

	user, err := u.userRepository.Create(newUser)

	if err != nil {
		if stdErrors.Is(err, errors.ErrDuplicateEntry) {
			return nil, errors.ErrUserAlreadyExists
		}
		return nil, err
	}

	return user, nil
}

func (u *UserService) Login(loginUserInput dto.LoginUserInput) (*string, error) {
	user, err := u.userRepository.FindByEmail(loginUserInput.Email)

	if err != nil {
		if stdErrors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	if err := utils.CheckPassword(user.PasswordHash, loginUserInput.Passowrd); err != nil {
		return nil, errors.ErrUserNotFound
	}

	token, err := utils.CreateToken(user)

	if err != nil {
		return nil, errors.ErrCreateToken
	}

	return token, nil
}

func (u *UserService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 型アサーション：SigningMethodHMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrParseToken
		}
		return []byte(config.GetEnv("SECRET_KEY", "")), nil
	})

	if err != nil {
		return nil, errors.ErrParseToken
	}

	// 型アサーション：MapClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 有効期限の確認
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.ErrExpiredToken
		}

		user, err := u.userRepository.FindByEmail(claims["email"].(string))

		if err != nil {
			return nil, errors.ErrUserNotFound
		}

		return user, nil
	}

	return nil, errors.ErrParseToken
}
