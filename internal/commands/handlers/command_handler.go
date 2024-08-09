package handlers

type ICommandHandler interface {
	ValidateArgs() bool
	HandleCommand()
}
