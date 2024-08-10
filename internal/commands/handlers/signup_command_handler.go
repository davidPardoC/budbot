package handlers

import (
	"log"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type SignupCommandHandler struct {
	telegramService services.ITelegramService
}

func NewSignupCommandHandler(telegramService services.ITelegramService) SignupCommandHandler {
	return SignupCommandHandler{telegramService: telegramService}
}

func (h SignupCommandHandler) HandleCommand(chatID int64) {
	log.Println("Handling signup command")
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	telegramMessageBuilder.SetText("Please provide your contact information")
}

func (h SignupCommandHandler) ValidateArgs() bool {
	return true
}
