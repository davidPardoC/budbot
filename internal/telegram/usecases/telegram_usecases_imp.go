package usecases

import (
	"fmt"
	"log"
	"slices"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/commands/factory"
	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
)

type TelegramUsecases struct {
	userUseCases userUc.IUserUseCases
	config       config.Config
	services     services.ITelegramService
}

func NewTelegramUsecases(userUseCases userUc.IUserUseCases, config config.Config, services services.ITelegramService) *TelegramUsecases {
	return &TelegramUsecases{
		userUseCases: userUseCases,
		config:       config,
		services:     services,
	}
}

func (u *TelegramUsecases) HandleWebhook(body dtos.TelegramWebhookDto) (string, error) {

	chatId := body.Message.Chat.Id

	if body.Message.ContactDto.UserId != 0 {
		user, err := u.userUseCases.CreateUser(body.Message.ContactDto.UserId, body.Message.ContactDto.PhoneNumber, body.Message.ContactDto.FirstName, body.Message.ContactDto.LastName, body.Message.Chat.Type)
		if err != nil {
			log.Println("Error creating user: ", err)
			return "pong", err
		}
		u.SendOnSignupMessage(user.ID)
		return "pong", nil
	}

	commandsFactory := factory.NewCommandsFactory(u.config, u.services)
	commands := commandsFactory.GetCommandsList()
	callbackQueryCommand := u.GetCommandFromCallbackQuery(body)

	user, err := u.userUseCases.FindByChatID(body.Message.Chat.Id)

	fmt.Println("Commands: ", commands)
	fmt.Println("CallbackQueryCommand: ", callbackQueryCommand)

	isCommand := slices.Contains(commands, callbackQueryCommand)
	handler := commandsFactory.GetCommand(callbackQueryCommand)

	if user == nil && !isCommand {
		log.Println("User does not exist")
		u.RequestForContact(chatId)
		return "pong", nil
	} else if isCommand && handler != nil && user == nil {
		log.Println("Command exists")
		handler.HandleCommand(chatId)
		return "pong", nil
	}

	if err != nil {
		return "", err
	}

	if user != nil {
		fmt.Println("User exists")
		return "pong", nil
	}

	return "pong", nil
}

func (u *TelegramUsecases) GetCommandFromCallbackQuery(webhookData dtos.TelegramWebhookDto) string {
	return webhookData.CallbackQuery.Data
}

func (u *TelegramUsecases) RequestForContact(chatID int64) (string, error) {

	mainText := `Seems like you are new here. Please provide your contact information by clicking the button below.`

	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatID)
	telegramMessageBuilder.SetText(mainText).SetParseMode("Markdown").AddKeyboardButton("Send my contact information", true)
	payload := telegramMessageBuilder.Build()
	u.services.SendMessage(payload)

	return "", nil
}

func (u *TelegramUsecases) SendOnSignupMessage(chatId int64) {
	mainText := `Welcome to BudBot!`
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatId)
	telegramMessageBuilder.SetText(mainText).SetParseMode("Markdown")
	payload := telegramMessageBuilder.Build()
	u.services.SendMessage(payload)
}
