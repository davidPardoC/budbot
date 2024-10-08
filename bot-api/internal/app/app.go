package app

import (
	"fmt"

	auth "github.com/davidPardoC/budbot/internal/auth/delivery/http"
	frontend "github.com/davidPardoC/budbot/internal/frontend/delivery/http"
	health "github.com/davidPardoC/budbot/internal/health/delivery/http"
	telegram "github.com/davidPardoC/budbot/internal/telegram/delivery/http"
	users "github.com/davidPardoC/budbot/internal/users/delivery/http"

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
	authRouter := auth.NewAuthRouter(a.gin, a.db, a.cfg)
	userRouter := users.NewUserRouter(a.gin, a.db)

	webhookRouter.SetupWebhookRouter()
	healthRouter.SetupHelthRouter()
	authRouter.SetupRoutes()
	userRouter.SetupRoutes()
	frontendRouter.SetupFrontendRouter()

	a.gin.Run(fmt.Sprintf(":%s", a.cfg.Server.Port))
}
