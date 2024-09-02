package handlers

import (
	"strconv"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"
)

type BudgetCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
	budget          float64
}

func NewBudgetCommandHandler(telegramService services.ITelegramService, userUseCases usecases.IUserUseCases) *BudgetCommandHandler {
	return &BudgetCommandHandler{
		telegramService: telegramService,
		userUseCases:    userUseCases,
	}
}

func (h *BudgetCommandHandler) HandleCommand(chatID int64, args []string) {

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)

	if !h.ValidateArgs(args) {
		payload := telegramMessageBuilder.SetText(messages.BudgetCommandInvalidArgsText).Build()
		h.telegramService.SendMessage(payload)
		return
	}

	message, _ := h.userUseCases.SetCurrentMothBudget(chatID, h.budget)

	payload := telegramMessageBuilder.SetText(message).Build()
	h.telegramService.SendMessage(payload)
}

func (h *BudgetCommandHandler) ValidateArgs(args []string) bool {
	if len(args) < 1 {
		return false
	}

	parsedAmount, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		return false
	}
	h.budget = parsedAmount
	return true
}
