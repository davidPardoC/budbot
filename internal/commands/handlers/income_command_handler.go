package handlers

import (
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"
)

type IncomeCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
}

func NewIncomeCommandHandler(telegramService services.ITelegramService, userUseCases usecases.IUserUseCases) *IncomeCommandHandler {
	return &IncomeCommandHandler{
		telegramService: telegramService,
		userUseCases:    userUseCases,
	}
}

func (h *IncomeCommandHandler) HandleCommand(chatID int64, args []string) {
}

func (h *IncomeCommandHandler) ValidateArgs(args []string) bool {
	return true
}
