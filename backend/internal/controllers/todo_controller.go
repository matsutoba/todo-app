package controllers

import (
	"net/http"
	"todo-app/internal/errors"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type ITodoController interface {
	FindAll(c *gin.Context)
}

type TodoController struct {
	service services.ITodoService
}

func NewTodoController(service services.ITodoService) ITodoController {
	return &TodoController{service: service}
}

func (t *TodoController) FindAll(c *gin.Context) {
	todos, err := t.service.FindAll()
	if err != nil {
		errors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}
