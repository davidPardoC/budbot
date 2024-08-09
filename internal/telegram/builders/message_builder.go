package builders

import (
	"encoding/json"
	"fmt"
)

type ParseMode string

const (
	Markdown ParseMode = "Markdown"
	HTML     ParseMode = "HTML"
)

type TelegramMessageBuilder struct {
	ChatID      int64       `json:"chat_id"`
	ParseMode   ParseMode   `json:"parse_mode"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
}

func NewTelegramMessageBuilder() *TelegramMessageBuilder {
	return &TelegramMessageBuilder{}
}

func (b *TelegramMessageBuilder) SetChatID(chatID int64) {
	b.ChatID = chatID
}

func (b *TelegramMessageBuilder) SetParseMode(parseMode ParseMode) {
	b.ParseMode = parseMode
}

func (b *TelegramMessageBuilder) SetText(text string) {
	b.Text = text
}

func (b *TelegramMessageBuilder) AddInlineKeyboardButton(text string) {
	b.ReplyMarkup.InlineKeyboard = append(b.ReplyMarkup.InlineKeyboard, InlineKeyboardButton{Text: text})
}

func (b *TelegramMessageBuilder) Build() string {
	encoded, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}
	return string(encoded)
}
