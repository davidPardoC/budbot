package handlers

type HelpCommandHandler struct{}

func (h HelpCommandHandler) HandleCommand() {
}

func (h HelpCommandHandler) ValidateArgs() bool {
	return true
}
