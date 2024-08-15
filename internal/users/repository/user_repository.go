package repository

import "github.com/davidPardoC/budbot/internal/users/models"

type IUserRepository interface {
	CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string) (models.User, error)
	FindByChatID(chatId int64) (models.User, error)
}
