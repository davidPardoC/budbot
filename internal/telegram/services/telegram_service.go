package services

type ITelegramService interface {
	SendMessage(payload string) error
}
