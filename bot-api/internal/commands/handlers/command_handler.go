package handlers

type ICommandHandler interface {
	ValidateArgs(args []string) bool
	HandleCommand(chatID int64, args []string)
}
