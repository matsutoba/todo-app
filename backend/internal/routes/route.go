package routes

import (
	"fmt"
	"log"
	"todo-app/internal/config"
	"todo-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() {
	host := config.GetEnv("API_HOST", "localhost")
	port := config.GetEnv("API_PORT", "8080")
	endpoint := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	r.POST("/register", handlers.RegisterUser)

	log.Print("Server running on ", endpoint)
	r.Run(endpoint)
}
