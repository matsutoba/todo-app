package controllers

import (
	"net/http"
	"todo-app/internal/dto"
	"todo-app/internal/errors"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}

func (u *UserController) Create(c *gin.Context) {
	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := u.service.Create(input)
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

func (u *UserController) Login(c *gin.Context) {
	var input dto.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := u.service.Login(input)
	if err != nil {
		errors.HandleError(c, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
