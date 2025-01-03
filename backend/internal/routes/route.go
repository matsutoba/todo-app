package routes

import (
	"fmt"
	"log"
	"todo-app/internal/config"
	"todo-app/internal/controllers"
	"todo-app/internal/middlewares"
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

	userRouter := r.Group("/users")
	userRouter.POST("/register", userController.Create)
	userRouter.POST("/login", userController.Login)

	todoRouter := r.Group("/todos", middlewares.AuthMiddleware(userService))
	todoRouter.GET("", todoController.FindAll)
	todoRouter.GET("/:id", todoController.FindById)
	todoRouter.POST("", todoController.Create)
	todoRouter.PUT("/:id", todoController.Update)
	todoRouter.DELETE("/:id", todoController.Delete)

	log.Print("Server running on ", endpoint)

	r.Run(endpoint)
}
