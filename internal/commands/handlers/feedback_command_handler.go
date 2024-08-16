package handlers

import (
	"strings"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type FeedBackCommandHandler struct {
	telegramService services.ITelegramService
}

func NewFeedBackCommandHandler(telegramService services.ITelegramService) *FeedBackCommandHandler {
	return &FeedBackCommandHandler{telegramService: telegramService}
}

func (h FeedBackCommandHandler) HandleCommand(chatID int64, args []string) {
	isValid := h.ValidateArgs(args)

	if !isValid {
		telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
		payload := telegramMessageBuilder.SetText(messages.FeedbackCommandInvalidArgsText).Build()
		h.telegramService.SendMessage(payload)
		return
	}

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	payload := telegramMessageBuilder.SetText(messages.FeedbackCommandText).Build()
	h.telegramService.SendMessage(payload)
}

func (h FeedBackCommandHandler) ValidateArgs(args []string) bool {
	jointArgs := strings.Join(args, " ")

	return jointArgs != ""
}
