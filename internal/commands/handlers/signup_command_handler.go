package handlers

type SignupCommandHandler struct{}

func NewSignupCommandHandler() SignupCommandHandler {
	return SignupCommandHandler{}
}

func (h SignupCommandHandler) HandleCommand() {
}

func (h SignupCommandHandler) ValidateArgs() bool {
	return true
}
