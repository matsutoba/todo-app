package controllers

import (
	"net/http"
	"strconv"
	"todo-app/internal/dto"
	"todo-app/internal/errors"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
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

func (t *TodoController) FindById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.ErrInvalidRequest)
		return
	}

	todo, err := t.service.FindById(uint(id))
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (t *TodoController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.ErrInvalidRequest)
		return
	}

	var input dto.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.HandleError(c, errors.ErrInvalidRequest)
		return
	}

	todo, err := t.service.Update(uint(id), input)
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}
