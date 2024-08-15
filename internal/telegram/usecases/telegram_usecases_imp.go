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

	user, _ := u.userUseCases.FindByChatID(body.Message.Chat.Id)

	chatId := body.Message.Chat.Id

	if body.Message.ContactDto.UserId != 0 && user == nil {
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

	fmt.Println("Commands: ", commands)
	fmt.Println("CallbackQueryCommand: ", callbackQueryCommand)

	isCommand := slices.Contains(commands, callbackQueryCommand)
	handler := commandsFactory.GetCommand(callbackQueryCommand)

	if user == nil && !isCommand {
		log.Println("User does not exist")
		u.RequestForContact(chatId)
		return "pong", nil
	}

	if user != nil && !isCommand {
		u.SendOnSignupMessage(user.ID)
		return "pong", nil
	}

	if isCommand {
		handler.HandleCommand(chatId)
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

	commands := []struct {
		Label   string
		Command string
	}{
		{"Get a list of commands", "/help"},
		{"Record a new category", "/rc <category_name>"},
		{"Record a new expense", "/re <category_name> <desc> <amount>"},
		{"Get a report of a category", "/report <category_name>"},
		{"Get a report of all categories", "/report"},
		{"Delete a category", "/dc <category_name>"},
	}

	mainText := `
	Welcome to BudBot!
	We are glad you are here. 
	Here is a list of commands you can use:
	`
	telegramMessageBuilder := builders.NewTelegramMessageBuilder(chatId)
	telegramMessageBuilder.SetText(mainText).SetParseMode("Markdown").RemovePreviousKeyboard()

	for _, command := range commands {
		telegramMessageBuilder.AddInlineKeyboardButton(command.Label, command.Command)
	}

	payload := telegramMessageBuilder.Build()
	u.services.SendMessage(payload)
}
