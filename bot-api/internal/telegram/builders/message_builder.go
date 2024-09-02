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

type TelegramMessage struct {
	ChatID          int64       `json:"chat_id"`
	ParseMode       ParseMode   `json:"parse_mode"`
	Text            string      `json:"text"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`
	ResizeKeyboard  bool        `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool        `json:"one_time_keyboard,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
	Keyboard       [][]KeyboardButton       `json:"keyboard,omitempty"`
	RemoveKeyboard bool                     `json:"remove_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

type KeyboardButton struct {
	Text           string `json:"text"`
	RequestContact bool   `json:"request_contact,omitempty"`
}

func NewTelegramMessageBuilder(chatID int64) *TelegramMessage {
	return &TelegramMessage{ChatID: chatID, OneTimeKeyboard: true, ResizeKeyboard: true}
}

func (b *TelegramMessage) SetChatID(chatID int64) *TelegramMessage {
	b.ChatID = chatID
	return b
}

func (b *TelegramMessage) SetParseMode(parseMode ParseMode) *TelegramMessage {
	b.ParseMode = parseMode
	return b
}

func (b *TelegramMessage) SetText(text string) *TelegramMessage {
	b.Text = text
	return b
}

func (b *TelegramMessage) AddInlineKeyboardButton(label string, callback string) *TelegramMessage {
	b.ReplyMarkup.InlineKeyboard = append(b.ReplyMarkup.InlineKeyboard, []InlineKeyboardButton{{Text: label, CallbackData: callback}})
	return b
}

func (b *TelegramMessage) AddKeyboardButton(text string, requestContact bool) *TelegramMessage {
	b.ReplyMarkup.Keyboard = append(b.ReplyMarkup.Keyboard, []KeyboardButton{{Text: text, RequestContact: requestContact}})
	return b
}

func (b *TelegramMessage) RemovePreviousKeyboard() *TelegramMessage {
	b.ReplyMarkup.RemoveKeyboard = true
	return b
}

func (b *TelegramMessage) Build() string {
	encoded, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}
	return string(encoded)
}
