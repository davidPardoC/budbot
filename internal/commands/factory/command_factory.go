package factory

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/commands/handlers"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type commandsFactory struct {
	config       config.Config
	commands     map[string]handlers.ICommandHandler
	commandsList []string
}

func NewCommandsFactory(config config.Config, telegramService services.ITelegramService) *commandsFactory {
	commands := make(map[string]handlers.ICommandHandler)

	commands["/signup"] = handlers.NewSignupCommandHandler(telegramService)

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
