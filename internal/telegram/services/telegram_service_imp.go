package services

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/davidPardoC/budbot/config"
)

type TelegramService struct {
	config config.Config
}

func NewTelegramService(config config.Config) ITelegramService {
	return &TelegramService{config: config}
}

func (ts *TelegramService) SendMessage(payload string) error {
	requestBody := bytes.NewBuffer([]byte(payload))
	url := fmt.Sprintf("%s/sendMessage", ts.config.Telegram.BaseURL)
	resp, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
