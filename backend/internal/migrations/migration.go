package migrations

import (
	"log"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database!")
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		panic("Failed to migrate database!")
	}

	log.Print("Migration has been processed")
}
