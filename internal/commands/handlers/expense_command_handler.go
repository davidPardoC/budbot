package handlers

import (
	"strconv"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"
)

type ExpenseCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
	amount          float64
	description     string
	category        string
}

func NewExpenseCommandHandler(telegramService services.ITelegramService, userUseCases usecases.IUserUseCases) *ExpenseCommandHandler {
	return &ExpenseCommandHandler{
		telegramService: telegramService,
		userUseCases:    userUseCases,
	}
}

func (h *ExpenseCommandHandler) HandleCommand(chatID int64, args []string) {

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)

	if !h.ValidateArgs(args) {
		payload := telegramMessageBuilder.SetText(messages.ExpenseCommandInvalidArgsText).Build()
		h.telegramService.SendMessage(payload)
		return
	}
}

func (h *ExpenseCommandHandler) ValidateArgs(args []string) bool {

	if len(args) < 1 {
		return false
	}

	amount := args[0]

	if args[1] != "" || args[2] != "" {
		h.description = args[1]
		h.category = args[2]
	}

	parsedAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return false
	}
	h.amount = parsedAmount

	return true
}
