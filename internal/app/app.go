package app

import (
	"fmt"

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
	webhookRouter := telegram.NewWebhookRouter(a.gin, a.cfg)
	healthRouter := health.NewHealthRouter(a.gin)

	webhookRouter.SetupWebhookRouter()
	healthRouter.SetupHelthRouter()

	a.gin.Run(fmt.Sprintf(":%s", a.cfg.Server.Port))
}
