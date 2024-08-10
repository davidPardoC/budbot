package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/davidPardoC/budbot/internal/telegram/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/telegram/usecases"
	"github.com/gin-gonic/gin"
)

type TelegramHandlers struct {
	usecases usecases.ITelegramUsecases
}

func NewTelegramHandlers(usecases usecases.ITelegramUsecases) *TelegramHandlers {
	return &TelegramHandlers{
		usecases: usecases,
	}
}

func (h *TelegramHandlers) WebHookHandler(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)

	var prettyBody bytes.Buffer
	if err := json.Indent(&prettyBody, body, "", "  "); err != nil {
		fmt.Println("Error formateando JSON:", err)
	} else {
		fmt.Println("Cuerpo de la solicitud:\n", prettyBody.String())
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var webhookBody dtos.TelegramWebhookDto

	if err := c.ShouldBindJSON(&webhookBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.usecases.HandleWebhook(webhookBody)

	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}
