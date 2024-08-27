package http

import (
	"fmt"
	"net/http"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/constants"
	"github.com/davidPardoC/budbot/internal/auth/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/auth/usecases"
	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	authUseCases usecases.IAuthUseCases
	config       config.Config
}

func NewAuthHandlers(authUseCases usecases.IAuthUseCases, config config.Config) *AuthHandlers {
	return &AuthHandlers{authUseCases: authUseCases, config: config}
}

func (a *AuthHandlers) oAuthTelegramHandler(c *gin.Context) {

	redirectDomain := constants.Domain

	if a.config.Server.Env == "local" {
		redirectDomain = constants.LocalDomain
	}

	callbackDto := dtos.TelegramCallbackDto{
		Id:        c.Query("id"),
		FirstName: c.Query("first_name"),
		LastName:  c.Query("last_name"),
		Username:  c.Query("username"),
		PhotoUrl:  c.Query("photo_url"),
		AuthDate:  c.Query("auth_date"),
		Hash:      c.Query("hash"),
	}

	credentials, err := a.authUseCases.Login(callbackDto, c.Request.URL.Query())

	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/error-login")
		return
	}

	c.SetCookie("token", credentials.Token, constants.HourInSeconds, "/", redirectDomain, false, false)
	c.SetCookie("refresh-token", credentials.RefreshToken, constants.WeekInSeconds, "/", redirectDomain, false, false)
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/home", redirectDomain))
}
