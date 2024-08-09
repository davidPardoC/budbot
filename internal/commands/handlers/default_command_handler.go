package handlers

type DefaultCommandHandler struct{}

func NewDefaultCommandHandler() DefaultCommandHandler {
	return DefaultCommandHandler{}
}

func (h DefaultCommandHandler) HandleCommand() {
}

func (h DefaultCommandHandler) ValidateArgs() bool {
	return true
}
