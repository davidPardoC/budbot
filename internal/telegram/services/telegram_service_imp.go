package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/davidPardoC/budbot/config"
)

type TelegramService struct {
	config config.Config
}

func NewTelegramService(config config.Config) ITelegramService {
	return &TelegramService{config: config}
}

func (ts *TelegramService) SendMessage(payload any) error {
	jsonBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(jsonBody)
	url := fmt.Sprintf("%s/sendMessage", ts.config.Telegram.BaseURL)
	println(url)
	resp, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Println(string(body))

	return nil
}
