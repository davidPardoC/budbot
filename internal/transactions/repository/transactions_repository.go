package repository

import (
	"github.com/davidPardoC/budbot/internal/transactions/models"
	transactionModels "github.com/davidPardoC/budbot/internal/transactions/models"
)

type ITransactionsRepository interface {
	CreateTransaction(amount float64, description string, categoryID int64, transactionType transactionModels.TransactionType, createdBy int64) (int64, error)
	GetAllTransactions() ([]models.Transactions, error)
	DeleteTransaction(transactionID int64) error
}
