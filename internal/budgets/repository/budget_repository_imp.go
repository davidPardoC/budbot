package repository

import (
	"github.com/davidPardoC/budbot/internal/budgets/models"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	database *gorm.DB
}

func NewBudgetRepository(database *gorm.DB) *BudgetRepository {
	return &BudgetRepository{database: database}
}

func (br *BudgetRepository) CreateBudget(createdBy int64, budget float64) error {
	b := &models.Budget{
		Amount:    budget,
		CreatedBy: createdBy,
	}

	return br.database.Create(b).Error
}

func (br *BudgetRepository) GetLastBudget(userId int64) *models.Budget {
	var budget models.Budget
	br.database.Where("created_by = ?", userId).Last(&budget)
	return &budget
}
