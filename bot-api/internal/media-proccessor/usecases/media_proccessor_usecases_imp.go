package usecases

import (
	"fmt"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/media-proccessor/dtos"
	"github.com/davidPardoC/budbot/internal/media-proccessor/services"
	telegramServices "github.com/davidPardoC/budbot/internal/telegram/services"
)

type MediaProcessorUsecases struct {
	mediaServices    services.IMediaProcessorService
	telegramServices telegramServices.ITelegramService
	config           config.Config
}

func NewMediaProcessorUsecases(mediaServices services.IMediaProcessorService, telegramServices telegramServices.ITelegramService, config config.Config) IMediaProcessorUsecases {
	return &MediaProcessorUsecases{
		mediaServices:    mediaServices,
		telegramServices: telegramServices,
		config:           config,
	}
}

func (mpu *MediaProcessorUsecases) ProccesAudioMessage(fileId string, chatId int64) dtos.ProccesedAudioDto {
	file, err := mpu.telegramServices.GetFilePath(fileId)

	if err != nil {
		return dtos.ProccesedAudioDto{}
	}

	filePath := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", mpu.config.Telegram.Token, file.Result.FilePath)

	fmt.Println("File path: ", filePath)

	proccesedAudio := mpu.mediaServices.ProccesAudioMessage(filePath, chatId)

	return proccesedAudio
}
