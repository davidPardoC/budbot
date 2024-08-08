package http

import "github.com/gin-gonic/gin"

func WebhookRoute(r *gin.Engine) {
	r.POST("/api/v1/webhook/telegram", TelegramWebHookHandler)
}
