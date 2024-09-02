package http

import "github.com/gin-gonic/gin"

func HealtCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
