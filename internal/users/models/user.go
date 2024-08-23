package models

import (
	budgetsModel "github.com/davidPardoC/budbot/internal/budgets/models"
	transactiosModel "github.com/davidPardoC/budbot/internal/transactions/models"
)

type User struct {
	ID           int64                           `json:"id" gorm:"primary_key"`
	ChatID       int64                           `json:"chat_id" gorm:"unique,not null"`
	FirstName    string                          `json:"first_name"`
	LastName     string                          `json:"last_name"`
	PhoneNumber  string                          `json:"phone_number"`
	UserType     string                          `json:"user_type"`
	Budgets      []budgetsModel.Budget           `json:"budgets" gorm:"foreignKey:CreatedBy"`
	Transactions []transactiosModel.Transactions `json:"transactions" gorm:"foreignKey:CreatedBy"`
}
