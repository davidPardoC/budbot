package builders_test

import (
	"testing"

	"github.com/davidPardoC/budbot/internal/telegram/builders"
	"github.com/stretchr/testify/assert"
)

func Test_TelegramMessageBuilder(t *testing.T) {
	expected := `{
  "chat_id": 123456789,
  "parse_mode": "Markdown",
  "text": "Please provide your contact information",
  "reply_markup": {
    "inline_keyboard": [
      [
        {
          "text": "/signup",
          "callback_data": "/signup"
        }
      ],
      [
        {
          "text": "/login",
          "callback_data": "/login"
        }
      ]
    ]
  },
  "resize_keyboard": true,
  "one_time_keyboard": true
}`
	telegramMessage := builders.NewTelegramMessageBuilder(123456789).SetText("Please provide your contact information").SetParseMode("Markdown").AddInlineKeyboardButton("/signup").AddInlineKeyboardButton("/login").Build()
	assert.Equal(t, expected, telegramMessage)
}
