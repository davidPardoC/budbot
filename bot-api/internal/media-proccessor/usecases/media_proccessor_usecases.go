package usecases

import "github.com/davidPardoC/budbot/internal/media-proccessor/dtos"

type IMediaProcessorUsecases interface {
	ProccesAudioMessage(fileId string, chatId int64) dtos.ProccesedAudioDto
}
