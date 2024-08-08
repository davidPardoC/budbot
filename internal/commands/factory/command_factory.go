package commands

type ICommandHandler interface{}

type PrintCommandHandler struct{}

var CommandFactory map[string]ICommandHandler

func init() {
	CommandFactory = make(map[string]ICommandHandler)
	CommandFactory["print"] = PrintCommandHandler{}
}
