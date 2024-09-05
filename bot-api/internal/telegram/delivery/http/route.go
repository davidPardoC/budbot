package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/davidPardoC/budbot/config"
	budgetRepo "github.com/davidPardoC/budbot/internal/budgets/repository"
	mediaServices "github.com/davidPardoC/budbot/internal/media-proccessor/services"
	mediaUseCases "github.com/davidPardoC/budbot/internal/media-proccessor/usecases"
	"github.com/davidPardoC/budbot/internal/telegram/services"
	telegramUc "github.com/davidPardoC/budbot/internal/telegram/usecases"
	transactionRepo "github.com/davidPardoC/budbot/internal/transactions/repository"
	userRepo "github.com/davidPardoC/budbot/internal/users/repository"
	userUc "github.com/davidPardoC/budbot/internal/users/usecases"
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
	transactionRepository := transactionRepo.NewTransactionsRepository(r.db)
	userUc := userUc.NewUserUsecases(userRepository, budgetRepository, transactionRepository)

	telegramServices := services.NewTelegramService(r.config)
	mediaServices := mediaServices.NewMediaProcessorService(telegramServices, r.config)
	mediaUseCases := mediaUseCases.NewMediaProcessorUsecases(mediaServices, telegramServices, r.config)
	telegramUseCases := telegramUc.NewTelegramUsecases(userUc, r.config, telegramServices, mediaUseCases)
	telegramHandlers := NewTelegramHandlers(telegramUseCases)

	r.router.POST("/api/v1/webhook/telegram", telegramHandlers.WebHookHandler)
}
