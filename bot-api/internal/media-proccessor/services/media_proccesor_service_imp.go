package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/media-proccessor/dtos"
	"github.com/davidPardoC/budbot/internal/telegram/services"
)

type MediaProcessorService struct {
	telegramService services.ITelegramService
	config          config.Config
}

type proccessAudioPayload struct {
	FilePath string `json:"file_path"`
	ChatId   int64  `json:"chat_id"`
}

func NewMediaProcessorService(telegramService services.ITelegramService, config config.Config) *MediaProcessorService {
	return &MediaProcessorService{
		telegramService: telegramService,
		config:          config,
	}
}

func (m *MediaProcessorService) ProccesAudioMessage(filePath string, chatId int64) dtos.ProccesedAudioDto {
	payload := proccessAudioPayload{
		FilePath: filePath,
		ChatId:   chatId,
	}

	jsonPayload, _ := json.Marshal(payload)
	bodyReader := bytes.NewReader(jsonPayload)

	resp, err := http.Post("http://localhost:8002/api/media-proccesor/process-audio", "application/json", bodyReader)

	if err != nil {
		log.Println(err)
		return dtos.ProccesedAudioDto{}
	}

	defer resp.Body.Close()

	var proccesedAudio dtos.ProccesedAudioDto

	err = json.NewDecoder(resp.Body).Decode(&proccesedAudio)

	if err != nil {
		log.Println(err)
		return dtos.ProccesedAudioDto{}
	}

	return proccesedAudio
}
