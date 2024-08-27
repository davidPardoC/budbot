package repository

import (
	"github.com/davidPardoC/budbot/internal/transactions/models"
	"gorm.io/gorm"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}

func (tr *TransactionsRepository) CreateTransaction(amount float64, description string, transactionType models.TransactionType, createdBy int64) (int64, error) {
	transaction := models.Transactions{
		Amount:      amount,
		Description: description,
		Type:        transactionType,
		CreatedBy:   createdBy,
	}

	result := tr.db.Create(&transaction)

	return transaction.ID, result.Error
}

func (tr *TransactionsRepository) GetAllTransactions() ([]models.Transactions, error) {
	var transactions []models.Transactions
	tr.db.Find(&transactions)
	return transactions, nil
}

func (tr *TransactionsRepository) DeleteTransaction(transactionID int64) error {
	return tr.db.Where("id = ?", transactionID).Delete(&models.Transactions{}).Error
}

func (tr *TransactionsRepository) GetExpensesByMonth(userId int64, month int) (float64, error) {
	var expense float64
	result := tr.db.Table("transactions").Select("COALESCE(sum(amount), 0)").Where("created_by = ? AND type = ? AND EXTRACT(MONTH FROM created_at) = ?", userId, models.Expense, month).Scan(&expense)
	return expense, result.Error
}

func (tr *TransactionsRepository) GetIncomesByMonth(userId int64, month int) (float64, error) {
	var income float64
	result := tr.db.Table("transactions").Select("COALESCE(sum(amount), 0)").Where("created_by = ? AND type = ? AND EXTRACT(MONTH FROM created_at) = ?", userId, models.Income, month).Scan(&income)
	return income, result.Error
}

func (tr *TransactionsRepository) GetExpensesBetweenDates(userId int64, startDate string, endDate string) (float64, error) {
	var expense float64
	result := tr.db.Table("transactions").Select("COALESCE(sum(amount), 0)").Where("created_by = ? AND type = ? AND created_at BETWEEN ? AND ?", userId, models.Expense, startDate, endDate).Scan(&expense)
	return expense, result.Error
}

func (tr *TransactionsRepository) GetIncomesBetweenDates(userId int64, startDate string, endDate string) (float64, error) {
	var income float64
	result := tr.db.Table("transactions").Select("COALESCE(sum(amount), 0)").Where("created_by = ? AND type = ? AND created_at BETWEEN ? AND ?", userId, models.Income, startDate, endDate).Scan(&income)
	return income, result.Error
}
