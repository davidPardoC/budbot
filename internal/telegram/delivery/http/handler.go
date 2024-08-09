package http

import "github.com/gin-gonic/gin"

func WebHookHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
