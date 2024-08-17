package handlers

import (
	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type HelpCommandHandler struct {
	telegramService services.ITelegramService
}

func NewHelpCommandHandler(telegramService services.ITelegramService) *HelpCommandHandler {
	return &HelpCommandHandler{
		telegramService: telegramService,
	}
}

func (h HelpCommandHandler) HandleCommand(chatID int64, args []string) {
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	payload := telegramMessageBuilder.SetText(messages.CommandsListText).Build()
	h.telegramService.SendMessage(payload)
}

func (h HelpCommandHandler) ValidateArgs(args []string) bool {
	return true
}
