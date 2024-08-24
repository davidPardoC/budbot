package repository

import (
	"github.com/davidPardoC/budbot/internal/transactions/models"
	transactionModels "github.com/davidPardoC/budbot/internal/transactions/models"
)

type ITransactionsRepository interface {
	CreateTransaction(amount float64, description string, transactionType transactionModels.TransactionType, createdBy int64) (int64, error)
	GetAllTransactions() ([]models.Transactions, error)
	DeleteTransaction(transactionID int64) error
	GetExpensesByMonth(userId int64, month int) (float64, error)
	GetIncomesByMonth(userId int64, month int) (float64, error)
}
