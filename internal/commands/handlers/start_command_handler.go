package handlers

type StartCommandHandler struct{}

func NewStartCommandHandler() StartCommandHandler {
	return StartCommandHandler{}
}

func (h StartCommandHandler) HandleCommand() {
}

func (h StartCommandHandler) ValidateArgs() bool {
	return true
}
