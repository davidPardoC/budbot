package usecases

import (
	"time"

	budgetRepo "github.com/davidPardoC/budbot/internal/budgets/repository"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/users/models"
	"github.com/davidPardoC/budbot/internal/users/repository"
)

type UserUseCases struct {
	userRepository   repository.IUserRepository
	budgetRepository budgetRepo.IBudgetRepository
}

func NewUserUsecases(userRepository repository.IUserRepository, budgetRepository budgetRepo.IBudgetRepository) *UserUseCases {
	return &UserUseCases{
		userRepository:   userRepository,
		budgetRepository: budgetRepository,
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
