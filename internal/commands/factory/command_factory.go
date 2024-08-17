package factory

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/commands/handlers"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	"github.com/davidPardoC/budbot/internal/users/usecases"
)

type commandsFactory struct {
	config       config.Config
	commands     map[string]handlers.ICommandHandler
	commandsList []string
}

func NewCommandsFactory(config config.Config, telegramService services.ITelegramService, userUseCases usecases.IUserUseCases) *commandsFactory {
	commands := make(map[string]handlers.ICommandHandler)

	commands["/signup"] = handlers.NewSignupCommandHandler(telegramService)
	commands["/help"] = handlers.NewHelpCommandHandler(telegramService)
	commands["/about"] = handlers.NewAboutCommandHandler(telegramService)
	commands["/feedback"] = handlers.NewFeedBackCommandHandler(telegramService)
	commands["/budget"] = handlers.NewBudgetCommandHandler(telegramService, userUseCases)

	commandsList := make([]string, 0)

	for k := range commands {
		commandsList = append(commandsList, k)
	}

	return &commandsFactory{config: config, commands: commands, commandsList: commandsList}
}

func (cf *commandsFactory) GetCommand(command string) handlers.ICommandHandler {
	return cf.commands[command]
}

func (cf *commandsFactory) GetCommandsList() []string {
	return cf.commandsList
}
