package http

import (
	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	telegramUc "github.com/davidPardoC/budbot/internal/telegram/usecases"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
	"github.com/gin-gonic/gin"
)

type WebhookRouter struct {
	router *gin.Engine
	config config.Config
}

func NewWebhookRouter(router *gin.Engine, config config.Config) *WebhookRouter {
	return &WebhookRouter{router: router, config: config}
}

func (r *WebhookRouter) SetupWebhookRouter() {

	services := services.NewTelegramService(r.config)
	userUc := userUc.NewUserUsecases()
	telegramUseCases := telegramUc.NewTelegramUsecases(userUc, r.config, services)
	telegramHandlers := NewTelegramHandlers(telegramUseCases)

	r.router.POST("/api/v1/webhook/telegram", telegramHandlers.WebHookHandler)
}
