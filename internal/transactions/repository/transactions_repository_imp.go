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
	result := tr.db.Table("transactions").Select("sum(amount)").Where("created_by = ? AND type = ? AND EXTRACT(MONTH FROM created_at) = ?", userId, models.Expense, month).Scan(&expense)
	return expense, result.Error
}

func (tr *TransactionsRepository) GetIncomesByMonth(userId int64, month int) (float64, error) {
	var income float64
	result := tr.db.Table("transactions").Select("sum(amount)").Where("created_by = ? AND type = ? AND EXTRACT(MONTH FROM created_at) = ?", userId, models.Income, month).Scan(&income)
	return income, result.Error
}
