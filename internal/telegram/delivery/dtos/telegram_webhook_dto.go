package dtos

type TelegramWebhookDto struct {
	UpdateId int64      `json:"update_id"`
	Message  MessageDto `json:"message"`
}

type MessageDto struct {
	MessageId int64   `json:"message_id"`
	From      UserDto `json:"from"`
	Chat      ChatDto `json:"chat"`
	Date      int64   `json:"date"`
	Text      string  `json:"text"`
}

type UserDto struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type ChatDto struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}
