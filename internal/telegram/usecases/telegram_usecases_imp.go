package usecases

import (
	"fmt"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
)

type TelegramUsecases struct {
	userUseCases userUc.IUserUseCases
	config       config.Config
}

func NewTelegramUsecases(userUseCases userUc.IUserUseCases, config config.Config) *TelegramUsecases {
	return &TelegramUsecases{
		userUseCases: userUseCases,
		config:       config,
	}
}

func (u *TelegramUsecases) HandleWebhook(body dtos.TelegramWebhookDto) (string, error) {
	user, err := u.userUseCases.FindByChatID(body.Message.Chat.Id)

	if err != nil {
		return "", err
	}

	if user != nil {
		fmt.Println("User exists")
		return "pong", nil
	} else {
		fmt.Println("User does not exist")
		u.RequestForContact(body.Message.Chat.Id)
		return "pong", nil
	}
}

func (u *TelegramUsecases) RequestForContact(chatID int64) (string, error) {

	telegramServices := services.NewTelegramService(u.config)

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	payload := telegramMessageBuilder.SetText("Please provide your contact information").SetParseMode("Markdown").AddInlineKeyboardButton("/signup").Build()

	telegramServices.SendMessage(payload)

	return "", nil
}
