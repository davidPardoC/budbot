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

func NewTelegramMessageBuilder(chatID int64) *TelegramMessageBuilder {
	return &TelegramMessageBuilder{ChatID: chatID}
}

func (b *TelegramMessageBuilder) SetChatID(chatID int64) *TelegramMessageBuilder {
	b.ChatID = chatID
	return b
}

func (b *TelegramMessageBuilder) SetParseMode(parseMode ParseMode) *TelegramMessageBuilder {
	b.ParseMode = parseMode
	return b
}

func (b *TelegramMessageBuilder) SetText(text string) *TelegramMessageBuilder {
	b.Text = text
	return b
}

func (b *TelegramMessageBuilder) AddInlineKeyboardButton(text string) *TelegramMessageBuilder {
	b.ReplyMarkup.InlineKeyboard = append(b.ReplyMarkup.InlineKeyboard, InlineKeyboardButton{Text: text})
	return b
}

func (b *TelegramMessageBuilder) Build() string {
	encoded, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}
	return string(encoded)
}
