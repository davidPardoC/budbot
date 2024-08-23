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
