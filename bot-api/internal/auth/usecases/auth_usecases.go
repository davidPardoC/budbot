package usecases

import (
	"github.com/davidPardoC/budbot/internal/auth/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/auth/models"
)

type IAuthUseCases interface {
	Login(dto dtos.TelegramCallbackDto, query map[string][]string) (*models.Credentials, error)
}
