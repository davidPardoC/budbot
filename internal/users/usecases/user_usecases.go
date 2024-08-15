package usecases

import "github.com/davidPardoC/budbot/internal/users/models"

type IUserUseCases interface {
	CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string) (models.User, error)
	FindByChatID(chatID int64) (*models.User, error)
}
