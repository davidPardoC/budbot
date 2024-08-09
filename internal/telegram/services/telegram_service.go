package services

type ITelegramService interface {
	SendMessage(payload any) error
}
