package controllers

import (
	"net/http"
	"todo-app/internal/dto"
	"todo-app/internal/errors"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
}

type TodoController struct {
	service services.ITodoService
}

func NewTodoController(service services.ITodoService) ITodoController {
	return &TodoController{service: service}
}

func (t *TodoController) Create(c *gin.Context) {
	var input dto.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.HandleError(c, errors.ErrInvalidRequest)
		return
	}

	todo, err := t.service.Create(input)
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func (t *TodoController) FindAll(c *gin.Context) {
	todos, err := t.service.FindAll()
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}
