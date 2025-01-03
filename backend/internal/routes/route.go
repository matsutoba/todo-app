package routes

import (
	"fmt"
	"log"
	"todo-app/internal/config"
	"todo-app/internal/controllers"
	"todo-app/internal/repositories"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) {
	host := config.GetEnv("API_HOST", "localhost")
	port := config.GetEnv("API_PORT", "8080")
	endpoint := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	r.POST("/register", userController.Create)
	r.POST("/login", userController.Login)
	r.GET("/todos", todoController.FindAll)

	log.Print("Server running on ", endpoint)

	r.Run(endpoint)
}
