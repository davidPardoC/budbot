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
	GetExpensesBetweenDates(userId int64, startDate string, endDate string) (float64, error)
	GetIncomesBetweenDates(userId int64, startDate string, endDate string) (float64, error)
	GetTransactionsBetweenDates(userId int64, startDate string, endDate string) ([]models.Transactions, error)
	GetTransactionsGroupedByCategory(userId int64, startDate string, endDate string) ([]models.TransactionsGroupedByCategory, error)
}
