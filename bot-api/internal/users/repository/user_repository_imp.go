package repository

import (
	"github.com/davidPardoC/budbot/internal/users/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(userId int64, phone_number string, firstName string, lasName string, userType string, photoUrl string) (*models.User, error) {

	user := models.User{
		ID:          userId,
		ChatID:      userId,
		PhoneNumber: phone_number,
		FirstName:   firstName,
		LastName:    lasName,
		UserType:    userType,
		PhotoUrl:    photoUrl,
	}

	result := u.db.Create(&user)
	return &user, result.Error
}

func (u *UserRepository) FindByChatID(chatId int64) (*models.User, error) {
	var user models.User
	result := u.db.Where("chat_id = ?", chatId).First(&user)
	return &user, result.Error
}

func (u *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	result := u.db.Save(&user)
	return user, result.Error
}
