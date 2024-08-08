package http

import "github.com/gin-gonic/gin"

type TelegramHandlers struct {
}

func TelegramWebHookHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
