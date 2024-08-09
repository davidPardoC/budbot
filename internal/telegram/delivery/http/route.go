package http

import (
	"github.com/davidPardoC/budbot/config"
	telegramUc "github.com/davidPardoC/budbot/internal/telegram/usecases"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
	"github.com/gin-gonic/gin"
)

type WebhookRouter struct {
	router *gin.Engine
	config config.Config
}

func NewWebhookRouter(router *gin.Engine, config config.Config) *WebhookRouter {
	return &WebhookRouter{router: router}
}

func (r *WebhookRouter) SetupWebhookRouter() {

	userUc := userUc.NewUserUsecases()
	telegramUseCases := telegramUc.NewTelegramUsecases(userUc, r.config)
	telegramHandlers := NewTelegramHandlers(telegramUseCases)

	r.router.POST("/api/v1/webhook/telegram", telegramHandlers.WebHookHandler)
}
