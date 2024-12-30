package main

import (
	"log"
	"todo-app/internal/config"
	"todo-app/internal/routes"
)

func main() {
	log.Print("Server start.")
	config.LoadEnv()
	env := config.GetEnv("ENV", "")
	log.Print("Environment: ", env)
	db := config.SetupDatabase()
	log.Print("Database connected.")
	routes.RegisterRoutes(db)
}
