package main

import (
	"log"
	"todo-app/internal/config"
)

func main() {
	log.Print("Server start.")
	config.LoadEnv()
	env := config.GetEnv("ENV", "")
	log.Print("Environment: ", env)
	config.SetupDatabase()
	log.Print("Database connected.")
}
