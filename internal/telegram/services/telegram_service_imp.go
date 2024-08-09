package services

import (
	"bytes"
	"encoding/json"
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

func (ts *TelegramService) SendMessage(payload map[string]interface{}) error {
	jsonBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(jsonBody)
	resp, err := http.Post(ts.config.Telegram.BaseURL+"/sendMessage", "application/json", requestBody)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Println(string(body))

	return nil
}
