package services

import "github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"

type ITelegramService interface {
	SendMessage(payload string) error
	GetFilePath(fileId string) (*dtos.GetFileDto, error)
}
