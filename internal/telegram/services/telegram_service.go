package services

type ITelegramService interface {
	SendMessage(payload map[string]interface{}) error
}
