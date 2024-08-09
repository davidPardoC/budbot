package database

import (
	"fmt"
	"log"

	"github.com/davidPardoC/budbot/config"
	categories "github.com/davidPardoC/budbot/internal/categories/models"
	users "github.com/davidPardoC/budbot/internal/users/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{}, &categories.Category{})
}
