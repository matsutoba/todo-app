package config

import (
	"fmt"
	"todo-app/internal/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	env := GetEnv("ENV", "develop")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnv("DB_USER", "user"),
		GetEnv("DB_PASSOWRD", "password"),
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_PORT", "3306"),
		GetEnv("DB_NAME", "todoapp"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if env == "develop" {
		migrations.Migration(db)
	}

	return db
}
