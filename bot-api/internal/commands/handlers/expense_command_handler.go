package handlers

import (
	"strconv"
	"strings"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"

	transactionModels "github.com/davidPardoC/budbot/internal/transactions/models"
)

type ExpenseCommandHandler struct {
	telegramService services.ITelegramService
	userUseCases    usecases.IUserUseCases
	amount          float64
	description     string
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
	_, err := h.userUseCases.RegisterTransaction(h.amount, h.description, transactionModels.Expense, chatID)

	if err != nil {
		payload := telegramMessageBuilder.SetText(messages.ErrorRegisteringExpenseText).Build()
		h.telegramService.SendMessage(payload)
	}

	payload := builders.NewTelegramMessageBuilder(chatID).SetText(messages.SuccessfullyRegisteredExpenseText).Build()
	h.telegramService.SendMessage(payload)
}

func (h *ExpenseCommandHandler) ValidateArgs(args []string) bool {

	if len(args) < 2 {
		return false
	}

	amount := args[1]

	if strings.Contains(amount, ",") {
		amount = strings.ReplaceAll(amount, ",", ".")
	}

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
