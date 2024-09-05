package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"
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
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return err
	}

	sb := string(body)
	fmt.Println(sb)

	return nil
}

func (ts *TelegramService) GetFilePath(fileId string) (*dtos.GetFileDto, error) {
	url := fmt.Sprintf("%s/getFile?file_id=%s", ts.config.Telegram.BaseURL, fileId)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var data dtos.GetFileDto

	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	fmt.Println(data)

	return &data, nil
}
