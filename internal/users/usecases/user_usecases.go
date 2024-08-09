package usecases

import "github.com/davidPardoC/budbot/internal/users/models"

type IUserUseCases interface {
	CreateUser()
	FindByChatID(chatID int64) (*models.User, error)
}
