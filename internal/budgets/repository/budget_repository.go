package repository

import (
	"github.com/davidPardoC/budbot/internal/budgets/models"
)

type IBudgetRepository interface {
	CreateBudget(createdBy int64, budget float64) error
	GetLastBudget(userId int64) *models.Budget
	UpdateBudget(Id int64, budget float64) error
	GetBudgetByMoth(userId int64, month int) (*models.Budget, error)
}
