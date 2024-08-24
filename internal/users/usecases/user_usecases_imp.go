package usecases

import (
	"errors"
	"time"

	budgetRepo "github.com/davidPardoC/budbot/internal/budgets/repository"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	transactionsModels "github.com/davidPardoC/budbot/internal/transactions/models"
	transactionsRepo "github.com/davidPardoC/budbot/internal/transactions/repository"
	"github.com/davidPardoC/budbot/internal/users/models"
	"github.com/davidPardoC/budbot/internal/users/repository"
	"gorm.io/gorm"
)

type UserUseCases struct {
	userRepository   repository.IUserRepository
	budgetRepository budgetRepo.IBudgetRepository
	transactionsRepo transactionsRepo.ITransactionsRepository
}

func NewUserUsecases(userRepository repository.IUserRepository, budgetRepository budgetRepo.IBudgetRepository, transactionsRepo transactionsRepo.ITransactionsRepository) *UserUseCases {
	return &UserUseCases{
		userRepository:   userRepository,
		budgetRepository: budgetRepository,
		transactionsRepo: transactionsRepo,
	}
}

func (u *UserUseCases) CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string) (models.User, error) {
	return u.userRepository.CreateUser(userId, phone_number, firstName, lasName, userType)
}

func (u *UserUseCases) FindByChatID(chatID int64) (*models.User, error) {
	return u.userRepository.FindByChatID(chatID)
}

func (u *UserUseCases) SetCurrentMothBudget(userId int64, budget float64) (string, error) {
	lastBudget := u.budgetRepository.GetLastBudget(userId)

	if lastBudget == nil {
		err := u.budgetRepository.CreateBudget(userId, budget)
		return messages.SuccesBudgetCommandText, err
	}

	currentMonth := time.Now().Month()
	lastBudgetMonth := lastBudget.CreatedAt.Month()

	if currentMonth == lastBudgetMonth {
		err := u.budgetRepository.UpdateBudget(lastBudget.ID, budget)
		return messages.SuccessUpdatedBudgetText, err
	}
	err := u.budgetRepository.CreateBudget(userId, budget)
	return messages.SuccesBudgetCommandText, err
}

func (u *UserUseCases) RegisterTransaction(amount float64, description string, transactionType transactionsModels.TransactionType, userId int64) (int64, error) {
	return u.transactionsRepo.CreateTransaction(amount, description, transactionType, userId)
}

func (u *UserUseCases) GetCurrentMothStats(userId int64) (*models.UserStats, error) {
	currentMonth := time.Now().Month()

	expenses, err := u.transactionsRepo.GetExpensesByMonth(userId, int(currentMonth))

	if err != nil {
		return nil, err
	}

	incomes, err := u.transactionsRepo.GetIncomesByMonth(userId, int(currentMonth))

	if err != nil {
		return nil, err
	}

	budget, err := u.budgetRepository.GetBudgetByMoth(userId, int(currentMonth))

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	stats := models.UserStats{
		TotalIncome:     incomes,
		Spent:           expenses,
		Budget:          budget.Amount,
		SpentPercentage: (expenses / budget.Amount) * 100,
	}

	return &stats, err
}
