package dtos

type TelegramWebhookDto struct {
	UpdateId      int64            `json:"update_id"`
	Message       MessageDto       `json:"message"`
	CallbackQuery CallbackQueryDto `json:"callback_query"`
}

type MessageDto struct {
	MessageId  int64      `json:"message_id"`
	From       UserDto    `json:"from"`
	Chat       ChatDto    `json:"chat"`
	Date       int64      `json:"date"`
	Text       string     `json:"text"`
	ContactDto ContactDto `json:"contact"`
	Voice      VoiceDto   `json:"voice"`
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

type CallbackQueryDto struct {
	Id         string     `json:"id"`
	From       UserDto    `json:"from"`
	MessageDto MessageDto `json:"message"`
	Data       string     `json:"data"`
}

type ContactDto struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int64  `json:"user_id"`
}

type VoiceDto struct {
	Duration     int64  `json:"duration"`
	FileId       string `json:"file_id"`
	MimeType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
	FileUniqueId string `json:"file_unique_id"`
}
