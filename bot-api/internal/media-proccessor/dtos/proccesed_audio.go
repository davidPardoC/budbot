package dtos

type ProccesedAudioDto struct {
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	Category string  `json:"category"`
	ChatId   int64   `json:"chat_id"`
}
