package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func oAuthTelegramHandler(c *gin.Context) {

	id := c.Query("id")
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	username := c.Query("username")
	photoURL := c.Query("photo_url")
	authDate := c.Query("auth_date")
	hash := c.Query("hash")

	log := fmt.Sprintf("id: %s, first_name: %s, last_name: %s, username: %s, photo_url: %s, auth_date: %s, hash: %s", id, firstName, lastName, username, photoURL, authDate, hash)
	println(log)
	c.Redirect(http.StatusTemporaryRedirect, "/home")
}
