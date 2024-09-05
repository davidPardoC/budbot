package services

import "github.com/davidPardoC/budbot/internal/media-proccessor/dtos"

type IMediaProcessorService interface {
	ProccesAudioMessage(filePath string, chatId int64) dtos.ProccesedAudioDto
}
