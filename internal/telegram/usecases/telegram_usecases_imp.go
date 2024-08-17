package usecases

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/budgets/repository"
	"github.com/davidPardoC/budbot/internal/commands/factory"
	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/constants/messages"
	"github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
	"gorm.io/gorm"
)

type TelegramUsecases struct {
	userUseCases     userUc.IUserUseCases
	config           config.Config
	services         services.ITelegramService
	budgetRepository repository.IBudgetRepository
}

func NewTelegramUsecases(userUseCases userUc.IUserUseCases, config config.Config, services services.ITelegramService) *TelegramUsecases {
	return &TelegramUsecases{
		userUseCases: userUseCases,
		config:       config,
		services:     services,
	}
}

func (u *TelegramUsecases) HandleWebhook(body dtos.TelegramWebhookDto) (string, error) {

	user, err := u.userUseCases.FindByChatID(body.Message.Chat.Id)

	chatId := body.Message.Chat.Id

	if body.Message.ContactDto.UserId != 0 && user.ID == 0 {
		user, err := u.userUseCases.CreateUser(body.Message.ContactDto.UserId, body.Message.ContactDto.PhoneNumber, body.Message.ContactDto.FirstName, body.Message.ContactDto.LastName, body.Message.Chat.Type)
		if err != nil {
			log.Println("Error creating user: ", err)
			return "pong", err
		}
		u.SendOnSignupMessage(user.ID)
		return "pong", nil
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		u.RequestForContact(chatId)
		return "pong", nil
	}

	commandsFactory := factory.NewCommandsFactory(u.config, u.services, u.userUseCases)
	commands := commandsFactory.GetCommandsList()

	userCommand, args := u.getUserCommand(body)

	isCommand := u.isCommand(userCommand)
	isKnownCommand := u.isKnownCommand(userCommand, commands)

	if isCommand && !isKnownCommand {
		handler := commandsFactory.GetCommand("/help")
		handler.HandleCommand(chatId, args)
		return "pong", nil
	}

	fmt.Println("Commands -> ", commands)
	fmt.Println("UserCommand -> ", userCommand)

	handler := commandsFactory.GetCommand(userCommand)

	if isKnownCommand {
		handler.HandleCommand(chatId, args)
	} else {
		handler := commandsFactory.GetCommand("/help")
		handler.HandleCommand(chatId, []string{})
	}

	return "pong", nil
}

func (u *TelegramUsecases) isCommand(message string) bool {
	return strings.HasPrefix(message, "/")
}

func (u *TelegramUsecases) isKnownCommand(command string, commands []string) bool {
	return slices.Contains(commands, command)
}

func (u *TelegramUsecases) getUserCommand(webhookData dtos.TelegramWebhookDto) (string, []string) {
	message := webhookData.CallbackQuery.Data

	if message == "" {
		message = webhookData.Message.Text
	}

	command := strings.Split(message, " ")[0]
	args := strings.Split(message, " ")[1:]

	return command, args
}

func (u *TelegramUsecases) RequestForContact(chatID int64) {
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	telegramMessageBuilder.SetText(messages.ContactRequestText).SetParseMode("Markdown").AddKeyboardButton("Send my contact information", true)
	payload := telegramMessageBuilder.Build()
	u.services.SendMessage(payload)

}

func (u *TelegramUsecases) SendOnSignupMessage(chatId int64) {
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatId)
	telegramMessageBuilder.SetText(messages.CommandsListText).SetParseMode("Markdown").RemovePreviousKeyboard()
	payload := telegramMessageBuilder.Build()
	u.services.SendMessage(payload)
}
