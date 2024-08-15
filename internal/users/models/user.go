package models

import "github.com/davidPardoC/budbot/internal/categories/models"

type User struct {
	ID          int64             `json:"id" gorm:"primary_key"`
	ChatID      int64             `json:"chat_id" gorm:"unique,not null"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	PhoneNumber string            `json:"phone_number"`
	UserType    string            `json:"user_type"`
	Categories  []models.Category `json:"categories" gorm:"foreignKey:CreatedBy"`
}
