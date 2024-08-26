package app

import (
	"fmt"

	auth "github.com/davidPardoC/budbot/internal/auth/delivery/http"
	frontend "github.com/davidPardoC/budbot/internal/frontend/delivery/http"
	health "github.com/davidPardoC/budbot/internal/health/delivery/http"
	telegram "github.com/davidPardoC/budbot/internal/telegram/delivery/http"

	"github.com/davidPardoC/budbot/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db  *gorm.DB
	gin *gin.Engine
	cfg config.Config
}

func NewApp(db *gorm.DB, gin *gin.Engine, cfg config.Config) *App {
	return &App{
		db:  db,
		gin: gin,
		cfg: cfg,
	}
}

func (a *App) Run() {
	webhookRouter := telegram.NewWebhookRouter(a.gin, a.cfg, a.db)
	healthRouter := health.NewHealthRouter(a.gin)
	frontendRouter := frontend.NewFrontendRouter(a.gin)
	authRuter := auth.NewAuthRouter(a.gin, a.db, a.cfg)

	frontendRouter.SetupFrontendRouter()
	webhookRouter.SetupWebhookRouter()
	healthRouter.SetupHelthRouter()
	authRuter.SetupRoutes()
	frontendRouter.SetupFrontendRouter()

	a.gin.Run(fmt.Sprintf(":%s", a.cfg.Server.Port))
}
