package handlers

import (
	"strconv"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"

	transactionModels "github.com/davidPardoC/budbot/internal/transactions/models"
)

type IncomeCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
	amount          float64
	description     string
}

func NewIncomeCommandHandler(telegramService services.ITelegramService, userUseCases usecases.IUserUseCases) *IncomeCommandHandler {
	return &IncomeCommandHandler{
		telegramService: telegramService,
		userUseCases:    userUseCases,
	}
}

func (h *IncomeCommandHandler) HandleCommand(chatID int64, args []string) {

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)

	if !h.ValidateArgs(args) {
		payload := telegramMessageBuilder.SetText(messages.IncomeCommandInvalidArgsText).Build()
		h.telegramService.SendMessage(payload)
		return
	}
	_, err := h.userUseCases.RegisterTransaction(h.amount, h.description, transactionModels.Income, chatID)

	if err != nil {
		payload := telegramMessageBuilder.SetText(messages.ErrorRegisteringIncomeText).Build()
		h.telegramService.SendMessage(payload)
	}

	payload := builders.NewTelegramMessageBuilder(chatID).SetText(messages.SuccessfullyRegisteredIncomeText).Build()
	h.telegramService.SendMessage(payload)
}

func (h *IncomeCommandHandler) ValidateArgs(args []string) bool {

	if len(args) < 2 {
		return false
	}

	amount := args[1]

	if args[0] != "" {
		h.description = args[0]
	}

	parsedAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		return false
	}

	h.amount = parsedAmount

	return true
}
