package commands

import "github.com/davidPardoC/budbot/internal/commands/handlers"

var CommandFactory map[string]handlers.ICommandHandler
var CommandsList []string

func SetupCommands() {
	CommandFactory = make(map[string]handlers.ICommandHandler)
	CommandFactory["/start"] = handlers.StartCommandHandler{}
	CommandFactory["/help"] = handlers.HelpCommandHandler{}
	CommandFactory["/signup"] = handlers.SignupCommandHandler{}
}

func GenerateCommandsStringList() {
	for k := range CommandFactory {
		CommandsList = append(CommandsList, k)
	}
}

func init() {
	SetupCommands()
	GenerateCommandsStringList()
}
