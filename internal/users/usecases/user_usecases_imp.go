package usecases

import (
	"github.com/davidPardoC/budbot/internal/users/models"
	"github.com/davidPardoC/budbot/internal/users/repository"
)

type UserUseCases struct {
	userRepository repository.IUserRepository
}

func NewUserUsecases(userRepository repository.IUserRepository) *UserUseCases {
	return &UserUseCases{
		userRepository: userRepository,
	}
}

func (u *UserUseCases) CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string) (models.User, error) {
	return u.userRepository.CreateUser(userId, phone_number, firstName, lasName, userType)
}

func (u *UserUseCases) FindByChatID(chatID int64) (*models.User, error) {
	return u.userRepository.FindByChatID(chatID)
}
