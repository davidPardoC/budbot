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

func (br *BudgetRepository) UpdateBudget(Id int64, budget float64) error {
	b := &models.Budget{
		Amount: budget,
	}

	return br.database.Model(&models.Budget{}).Where("id = ?", Id).Updates(b).Error
}

func (br *BudgetRepository) GetBudgetByMoth(userId int64, month int) (*models.Budget, error) {
	var budget models.Budget
	result := br.database.Where("created_by = ? AND EXTRACT(MONTH FROM created_at) = ?", userId, month).First(&budget)
	return &budget, result.Error
}

func (br *BudgetRepository) GetBudgetBetweenDates(userId int64, startDate string, endDate string) (*models.Budget, error) {
	var budget models.Budget
	result := br.database.Where("created_by = ?", userId).Where("created_at >= ?", startDate).Where("created_at <= ?", endDate).First(&budget)
	return &budget, result.Error
}
