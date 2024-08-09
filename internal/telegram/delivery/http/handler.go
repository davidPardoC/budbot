package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func WebHookHandler(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error reading body",
		})
		return
	}

	var prettyJSON bytes.Buffer

	json.Indent(&prettyJSON, bodyBytes, "", "    ")

	fmt.Println(prettyJSON.String())

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
