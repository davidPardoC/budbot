package handlers

import (
	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type AboutCommandHandler struct {
	telegramService services.ITelegramService
}

func NewAboutCommandHandler(telegramService services.ITelegramService) *AboutCommandHandler {
	return &AboutCommandHandler{
		telegramService: telegramService,
	}
}

func (h AboutCommandHandler) HandleCommand(chatID int64, args []string) {
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	payload := telegramMessageBuilder.SetText(messages.AboutCommandText).Build()
	h.telegramService.SendMessage(payload)
}

func (h AboutCommandHandler) ValidateArgs(args []string) bool {
	return true
}
