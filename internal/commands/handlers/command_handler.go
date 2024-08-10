package handlers

type ICommandHandler interface {
	ValidateArgs() bool
	HandleCommand(chatID int64)
}
