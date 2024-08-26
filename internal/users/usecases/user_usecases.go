package usecases

import (
	transactionModels "github.com/davidPardoC/budbot/internal/transactions/models"
	"github.com/davidPardoC/budbot/internal/users/models"
)

type IUserUseCases interface {
	CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string) (*models.User, error)
	FindByChatID(chatID int64) (*models.User, error)
	SetCurrentMothBudget(userId int64, budget float64) (string, error)
	RegisterTransaction(amount float64, description string, transactionType transactionModels.TransactionType, userId int64) (int64, error)
	GetCurrentMothStats(userId int64) (*models.UserStats, error)
}
