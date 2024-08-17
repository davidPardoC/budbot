package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	telegramUc "github.com/davidPardoC/budbot/internal/telegram/usecases"
	userRepo "github.com/davidPardoC/budbot/internal/users/repository"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"

	budgetRepo "github.com/davidPardoC/budbot/internal/budgets/repository"
)

type WebhookRouter struct {
	router *gin.Engine
	config config.Config
	db     *gorm.DB
}

func NewWebhookRouter(router *gin.Engine, config config.Config, db *gorm.DB) *WebhookRouter {
	return &WebhookRouter{router: router, config: config, db: db}
}

func (r *WebhookRouter) SetupWebhookRouter() {

	budgetRepository := budgetRepo.NewBudgetRepository(r.db)
	userRepository := userRepo.NewUserRepository(r.db)
	userUc := userUc.NewUserUsecases(userRepository, budgetRepository)

	telegramServices := services.NewTelegramService(r.config)
	telegramUseCases := telegramUc.NewTelegramUsecases(userUc, r.config, telegramServices)
	telegramHandlers := NewTelegramHandlers(telegramUseCases)

	r.router.POST("/api/v1/webhook/telegram", telegramHandlers.WebHookHandler)
}
