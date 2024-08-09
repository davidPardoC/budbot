package usecases

import "github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"

type ITelegramUsecases interface {
	HandleWebhook(body dtos.TelegramWebhookDto) (string, error)
}
