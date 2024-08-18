package repository

import (
	"github.com/davidPardoC/budbot/internal/transactions/models"
	"gorm.io/gorm"
)

type TransactionsRepository struct {
	database *gorm.DB
}

func NewTransactionsRepository() *TransactionsRepository {
	return &TransactionsRepository{}
}

func (tr *TransactionsRepository) CreateTransaction(amount float64, description string, categoryID int64, transactionType models.TransactionType, createdBy int64) (int64, error) {
	t := &models.Transactions{
		Amount:      amount,
		Description: description,
		CategoryID:  categoryID,
		Type:        transactionType,
		CreatedBy:   createdBy,
	}

	return t.ID, tr.database.Create(t).Error
}

func (tr *TransactionsRepository) GetAllTransactions() ([]models.Transactions, error) {
	var transactions []models.Transactions
	tr.database.Find(&transactions)
	return transactions, nil
}

func (tr *TransactionsRepository) DeleteTransaction(transactionID int64) error {
	return tr.database.Where("id = ?", transactionID).Delete(&models.Transactions{}).Error
}
