package usecases

import "github.com/davidPardoC/budbot/internal/users/models"

type UserUseCases struct{}

func NewUserUsecases() *UserUseCases {
	return &UserUseCases{}
}

func (u *UserUseCases) CreateUser() {
}

func (u *UserUseCases) FindByChatID(chatID int64) (*models.User, error) {
	return nil, nil
}
