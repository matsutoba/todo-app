package utils

import (
	"time"
	"todo-app/internal/config"
	"todo-app/internal/models"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   jwt.TimeFunc().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetEnv("SECRET_KEY", "")))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
